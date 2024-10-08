package router

import (
	"net/http"

	"github.com/openmeterio/openmeter/api"
	creditdriver "github.com/openmeterio/openmeter/openmeter/credit/driver"
)

// List grants
// (GET /api/v1/grants)
func (a *Router) ListGrants(w http.ResponseWriter, r *http.Request, params api.ListGrantsParams) {
	if !a.config.EntitlementsEnabled {
		unimplemented.ListGrants(w, r, params)
		return
	}
	a.creditHandler.ListGrants().With(creditdriver.ListGrantsHandlerParams{
		Params: params,
	}).ServeHTTP(w, r)
}

// Delete a grant
// (DELETE /api/v1/grants/{grantId})
func (a *Router) VoidGrant(w http.ResponseWriter, r *http.Request, grantId api.GrantId) {
	if !a.config.EntitlementsEnabled {
		unimplemented.VoidGrant(w, r, grantId)
		return
	}
	a.creditHandler.VoidGrant().With(creditdriver.VoidGrantHandlerParams{
		ID: grantId,
	}).ServeHTTP(w, r)
}
