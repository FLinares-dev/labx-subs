package labx_subs

import (
	"context"
	"labxsubs/model"
)

type accessResolver struct{ *Resolver }

func (r *accessResolver) Supplier(ctx context.Context, obj *model.Access) (*Supplier, error) {
	panic("not implemented")
}
