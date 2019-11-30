package labx_subs

import (
	"context"
	"labxsubs/model"
)

type Resolver struct{}

func (r *Resolver) Access() AccessResolver {
	return &accessResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateAccess(ctx context.Context, input *AccessInput) (*model.Access, error) {
	panic("not implemented")
}
func (r *mutationResolver) UpdateAccess(ctx context.Context, id string, input *AccessInput) (*model.Access, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Accesses(ctx context.Context) ([]*model.Access, error) {
	panic("not implemented")
}
