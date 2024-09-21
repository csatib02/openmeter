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
	"fmt"
	"net/http"
	"time"

	"github.com/openmeterio/openmeter/api"
	"github.com/openmeterio/openmeter/openmeter/notification"
	"github.com/openmeterio/openmeter/pkg/defaultx"
	"github.com/openmeterio/openmeter/pkg/framework/commonhttp"
	"github.com/openmeterio/openmeter/pkg/framework/transport/httptransport"
	"github.com/openmeterio/openmeter/pkg/models"
	"github.com/openmeterio/openmeter/pkg/pagination"
	"github.com/openmeterio/openmeter/pkg/sortx"
)

type (
	ListEventsRequest  = notification.ListEventsInput
	ListEventsResponse = api.NotificationEventsResponse
	ListEventsParams   = api.ListNotificationEventsParams
	ListEventsHandler  httptransport.HandlerWithArgs[ListEventsRequest, ListEventsResponse, ListEventsParams]
)

func (h *handler) ListEvents() ListEventsHandler {
	return httptransport.NewHandlerWithArgs(
		func(ctx context.Context, r *http.Request, params ListEventsParams) (ListEventsRequest, error) {
			ns, err := h.resolveNamespace(ctx)
			if err != nil {
				return ListEventsRequest{}, fmt.Errorf("failed to resolve namespace: %w", err)
			}

			req := ListEventsRequest{
				Namespaces: []string{ns},
				Order:      sortx.Order(defaultx.WithDefault(params.Order, api.ListNotificationEventsParamsOrderSortOrderDESC)),
				OrderBy:    defaultx.WithDefault(params.OrderBy, notification.EventOrderByCreatedAt),
				Page: pagination.Page{
					PageSize:   defaultx.WithDefault(params.PageSize, notification.DefaultPageSize),
					PageNumber: defaultx.WithDefault(params.Page, notification.DefaultPageNumber),
				},
				Subjects: defaultx.WithDefault(params.Subject, nil),
				Features: defaultx.WithDefault(params.Feature, nil),
				Rules:    defaultx.WithDefault(params.Rule, nil),
				Channels: defaultx.WithDefault(params.Channel, nil),
				From:     defaultx.WithDefault(params.From, time.Time{}),
				To:       defaultx.WithDefault(params.To, time.Time{}),
			}

			return req, nil
		},
		func(ctx context.Context, request ListEventsRequest) (ListEventsResponse, error) {
			resp, err := h.service.ListEvents(ctx, request)
			if err != nil {
				return ListEventsResponse{}, fmt.Errorf("failed to list events: %w", err)
			}

			items := make([]api.NotificationEvent, 0, len(resp.Items))

			for _, event := range resp.Items {
				var item api.NotificationEvent

				item, err = event.AsNotificationEvent()
				if err != nil {
					return ListEventsResponse{}, fmt.Errorf("failed to cast event: %w", err)
				}

				items = append(items, item)
			}

			return ListEventsResponse{
				Items:      items,
				Page:       resp.Page.PageNumber,
				PageSize:   resp.Page.PageSize,
				TotalCount: resp.TotalCount,
			}, nil
		},
		commonhttp.JSONResponseEncoderWithStatus[ListEventsResponse](http.StatusOK),
		httptransport.AppendOptions(
			h.options,
			httptransport.WithOperationName("listNotificationEvents"),
			httptransport.WithErrorEncoder(errorEncoder()),
		)...,
	)
}

type (
	GetEventRequest  = notification.GetEventInput
	GetEventResponse = api.NotificationEvent
	GetEventHandler  httptransport.HandlerWithArgs[GetEventRequest, GetEventResponse, api.EventId]
)

func (h *handler) GetEvent() GetEventHandler {
	return httptransport.NewHandlerWithArgs(
		func(ctx context.Context, r *http.Request, eventID api.EventId) (GetEventRequest, error) {
			ns, err := h.resolveNamespace(ctx)
			if err != nil {
				return GetEventRequest{}, fmt.Errorf("failed to resolve namespace: %w", err)
			}

			req := GetEventRequest{
				NamespacedID: models.NamespacedID{
					Namespace: ns,
					ID:        eventID,
				},
			}

			return req, nil
		},
		func(ctx context.Context, request GetEventRequest) (GetEventResponse, error) {
			event, err := h.service.GetEvent(ctx, request)
			if err != nil {
				return GetEventResponse{}, fmt.Errorf("failed to get event: %w", err)
			}

			return event.AsNotificationEvent()
		},
		commonhttp.JSONResponseEncoderWithStatus[GetEventResponse](http.StatusOK),
		httptransport.AppendOptions(
			h.options,
			httptransport.WithOperationName("getNotificationEvent"),
			httptransport.WithErrorEncoder(errorEncoder()),
		)...,
	)
}
