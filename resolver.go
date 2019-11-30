package labx_subs

import (
	"context"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAccess(ctx context.Context, input *AccessInput) (*Access, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAccess(ctx context.Context, id string, input *AccessInput) (*Access, error) {
	panic("not implemented")
}
func (r *mutationResolver) CreateSupplier(ctx context.Context, input *SupplierInput) (*Supplier, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateSupplier(ctx context.Context, id string, input *SupplierInput) (*Supplier, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Accesses(ctx context.Context) ([]*Access, error) {
	panic("not implemented")
}
func (r *queryResolver) Suppliers(ctx context.Context) ([]*Supplier, error) {
	panic("not implemented")
}
