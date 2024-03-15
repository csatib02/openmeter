// Copyright © 2024 Tailfin Cloud Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httpdriver

import (
	"context"
	"errors"
	"net/http"
	"slices"

	"github.com/openmeterio/openmeter/api"
	"github.com/openmeterio/openmeter/internal/credit"
	entitlement_httpdriver "github.com/openmeterio/openmeter/internal/entitlement/httpdriver"
	meteredentitlement "github.com/openmeterio/openmeter/internal/entitlement/metered"
	"github.com/openmeterio/openmeter/internal/namespace/namespacedriver"
	"github.com/openmeterio/openmeter/pkg/convert"
	"github.com/openmeterio/openmeter/pkg/defaultx"
	"github.com/openmeterio/openmeter/pkg/framework/commonhttp"
	"github.com/openmeterio/openmeter/pkg/framework/transport/httptransport"
	"github.com/openmeterio/openmeter/pkg/models"
	"github.com/openmeterio/openmeter/pkg/pagination"
	"github.com/openmeterio/openmeter/pkg/strcase"
)

type GrantHandler interface {
	ListGrants() ListGrantsHandler
	VoidGrant() VoidGrantHandler
}

type grantHandler struct {
	namespaceDecoder namespacedriver.NamespaceDecoder
	options          []httptransport.HandlerOption
	grantConnector   credit.GrantConnector
}

func NewGrantHandler(
	namespaceDecoder namespacedriver.NamespaceDecoder,
	grantConnector credit.GrantConnector,
	options ...httptransport.HandlerOption,
) GrantHandler {
	return &grantHandler{
		namespaceDecoder: namespaceDecoder,
		grantConnector:   grantConnector,
		options:          options,
	}
}

type (
	ListGrantsHandlerRequest struct {
		params credit.ListGrantsParams
	}
	ListGrantsHandlerResponse = commonhttp.Union[[]api.EntitlementGrant, pagination.PagedResponse[api.EntitlementGrant]]
	ListGrantsHandlerParams   struct {
		Params api.ListGrantsParams
	}
)
type ListGrantsHandler httptransport.HandlerWithArgs[ListGrantsHandlerRequest, ListGrantsHandlerResponse, ListGrantsHandlerParams]

func (h *grantHandler) ListGrants() ListGrantsHandler {
	return httptransport.NewHandlerWithArgs[ListGrantsHandlerRequest, ListGrantsHandlerResponse, ListGrantsHandlerParams](
		func(ctx context.Context, r *http.Request, params ListGrantsHandlerParams) (ListGrantsHandlerRequest, error) {
			ns, err := h.resolveNamespace(ctx)
			if err != nil {
				return ListGrantsHandlerRequest{}, err
			}

			// validate OrderBy
			if params.Params.OrderBy != nil {
				if !slices.Contains(credit.GrantOrderBy("").StrValues(), strcase.CamelToSnake(string(*params.Params.OrderBy))) {
					return ListGrantsHandlerRequest{}, commonhttp.NewHTTPError(http.StatusBadRequest, errors.New("invalid order by"))
				}
			}

			handlerParams := credit.ListGrantsParams{
				Namespace:      ns,
				IncludeDeleted: defaultx.WithDefault(params.Params.IncludeDeleted, false),
				Page: pagination.Page{
					PageSize:   defaultx.WithDefault(params.Params.PageSize, 0),
					PageNumber: defaultx.WithDefault(params.Params.Page, 0),
				},
				Limit:  defaultx.WithDefault(params.Params.Limit, commonhttp.DefaultPageSize),
				Offset: defaultx.WithDefault(params.Params.Offset, 0),
				OrderBy: credit.GrantOrderBy(
					strcase.CamelToSnake(defaultx.WithDefault((*string)(params.Params.OrderBy), string(credit.GrantOrderByEffectiveAt))),
				),
				Order:            commonhttp.GetSortOrder(api.ListGrantsParamsOrderSortOrderASC, params.Params.Order),
				SubjectKeys:      convert.DerefHeaderPtr[string](params.Params.Subject),
				FeatureIdsOrKeys: convert.DerefHeaderPtr[string](params.Params.Feature),
			}
			if !handlerParams.Page.IsZero() {
				handlerParams.Page.PageNumber = defaultx.IfZero(handlerParams.Page.PageNumber, commonhttp.DefaultPage)
				handlerParams.Page.PageSize = defaultx.IfZero(handlerParams.Page.PageSize, commonhttp.DefaultPageSize)
			}
			return ListGrantsHandlerRequest{
				params: handlerParams,
			}, nil
		},
		func(ctx context.Context, request ListGrantsHandlerRequest) (ListGrantsHandlerResponse, error) {
			// due to backward compatibility, if pagination is not provided we return a simple array
			response := ListGrantsHandlerResponse{
				Option1: &[]api.EntitlementGrant{},
				Option2: &pagination.PagedResponse[api.EntitlementGrant]{},
			}
			grants, err := h.grantConnector.ListGrants(ctx, request.params)
			if err != nil {
				return response, err
			}

			apiGrants := make([]api.EntitlementGrant, 0, len(grants.Items))
			for _, grant := range grants.Items {
				entitlementGrant, err := meteredentitlement.GrantFromCreditGrant(grant)
				if err != nil {
					return response, err
				}
				// FIXME: not elegant but good for now, entitlement grants are all we have...
				apiGrant := entitlement_httpdriver.MapEntitlementGrantToAPI(nil, entitlementGrant)

				apiGrants = append(apiGrants, apiGrant)
			}

			if request.params.Page.IsZero() {
				response.Option1 = &apiGrants
			} else {
				response.Option1 = nil
				response.Option2 = &pagination.PagedResponse[api.EntitlementGrant]{
					Items:      apiGrants,
					TotalCount: grants.TotalCount,
					Page:       grants.Page,
				}
			}

			return response, nil
		},
		commonhttp.JSONResponseEncoder[ListGrantsHandlerResponse],
		httptransport.AppendOptions(
			h.options,
			httptransport.WithErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter, _ *http.Request) bool {
				if _, ok := err.(*models.GenericUserError); ok {
					commonhttp.NewHTTPError(
						http.StatusBadRequest,
						err,
					).EncodeError(ctx, w)
					return true
				}
				return commonhttp.HandleErrorIfTypeMatches[*pagination.InvalidError](ctx, http.StatusBadRequest, err, w)
			}),
		)...,
	)
}

type VoidGrantHandlerRequest struct {
	ID models.NamespacedID
}
type (
	VoidGrantHandlerResponse = interface{}
	VoidGrantHandlerParams   struct {
		ID string
	}
)
type VoidGrantHandler httptransport.HandlerWithArgs[VoidGrantHandlerRequest, VoidGrantHandlerResponse, VoidGrantHandlerParams]

func (h *grantHandler) VoidGrant() VoidGrantHandler {
	return httptransport.NewHandlerWithArgs[VoidGrantHandlerRequest, VoidGrantHandlerResponse, VoidGrantHandlerParams](
		func(ctx context.Context, r *http.Request, params VoidGrantHandlerParams) (VoidGrantHandlerRequest, error) {
			ns, err := h.resolveNamespace(ctx)
			if err != nil {
				return VoidGrantHandlerRequest{}, err
			}

			return VoidGrantHandlerRequest{
				ID: models.NamespacedID{
					Namespace: ns,
					ID:        params.ID,
				},
			}, nil
		},
		func(ctx context.Context, request VoidGrantHandlerRequest) (interface{}, error) {
			err := h.grantConnector.VoidGrant(ctx, request.ID)
			if err != nil {
				return nil, err
			}

			return nil, nil
		},
		commonhttp.EmptyResponseEncoder[interface{}](http.StatusNoContent),
		httptransport.AppendOptions(
			h.options,
			httptransport.WithErrorEncoder(func(ctx context.Context, err error, w http.ResponseWriter, _ *http.Request) bool {
				if _, ok := err.(*models.GenericUserError); ok {
					commonhttp.NewHTTPError(
						http.StatusBadRequest,
						err,
					).EncodeError(ctx, w)
					return true
				}
				if _, ok := err.(*credit.GrantNotFoundError); ok {
					commonhttp.NewHTTPError(
						http.StatusNotFound,
						err,
					).EncodeError(ctx, w)
					return true
				}
				return false
			}),
		)...,
	)
}

func (h *grantHandler) resolveNamespace(ctx context.Context) (string, error) {
	ns, ok := h.namespaceDecoder.GetNamespace(ctx)
	if !ok {
		return "", commonhttp.NewHTTPError(http.StatusInternalServerError, errors.New("internal server error"))
	}

	return ns, nil
}
