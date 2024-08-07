// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"

	"github.com/openmeterio/openmeter/pkg/pagination"
)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (bs *BalanceSnapshotQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*BalanceSnapshot], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	bs.ctx.Offset = &zero
	bs.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := bs.Clone()
	pagedQuery := bs

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*BalanceSnapshot]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*BalanceSnapshot] = (*BalanceSnapshotQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (e *EntitlementQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*Entitlement], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	e.ctx.Offset = &zero
	e.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := e.Clone()
	pagedQuery := e

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*Entitlement]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*Entitlement] = (*EntitlementQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (f *FeatureQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*Feature], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	f.ctx.Offset = &zero
	f.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := f.Clone()
	pagedQuery := f

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*Feature]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*Feature] = (*FeatureQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (gr *GrantQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*Grant], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	gr.ctx.Offset = &zero
	gr.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := gr.Clone()
	pagedQuery := gr

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*Grant]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*Grant] = (*GrantQuery)(nil)

// Paginate runs the query and returns a paginated response.
// If page is its 0 value then it will return all the items and populate the response page accordingly.
func (ur *UsageResetQuery) Paginate(ctx context.Context, page pagination.Page) (pagination.PagedResponse[*UsageReset], error) {
	// Get the limit and offset
	limit, offset := page.Limit(), page.Offset()

	// Unset previous pagination settings
	zero := 0
	ur.ctx.Offset = &zero
	ur.ctx.Limit = &zero

	// Create duplicate of the query to run for
	countQuery := ur.Clone()
	pagedQuery := ur

	// Unset ordering for count query
	countQuery.order = nil

	pagedResponse := pagination.PagedResponse[*UsageReset]{
		Page: page,
	}

	// Get the total count
	count, err := countQuery.Count(ctx)
	if err != nil {
		return pagedResponse, fmt.Errorf("failed to get count: %w", err)
	}
	pagedResponse.TotalCount = count

	// If page is its 0 value then return all the items
	if page.IsZero() {
		offset = 0
		limit = count
	}

	// Set the limit and offset
	pagedQuery.ctx.Limit = &limit
	pagedQuery.ctx.Offset = &offset

	// Get the paged items
	items, err := pagedQuery.All(ctx)
	pagedResponse.Items = items
	return pagedResponse, err
}

// type check
var _ pagination.Paginator[*UsageReset] = (*UsageResetQuery)(nil)
