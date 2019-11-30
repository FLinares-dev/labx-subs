package labx_subs

import (
	"context"
	"errors"
)

var mapSuppliers map[string]*Supplier

type supplierResolver struct{ *Resolver }

func (r *mutationResolver) CreateSupplier(ctx context.Context, input *SupplierInput) (*Supplier, error) {

	supplier := &Supplier{
		ID:   string(len(mapSuppliers)),
		Code: input.Code,
		Name: input.Name,
	}
	mapSuppliers[supplier.ID] = supplier
	return supplier, nil
}

func (r *mutationResolver) UpdateSupplier(ctx context.Context, id string, input *SupplierInput) (*Supplier, error) {
	supplier := &Supplier{
		ID:   id,
		Code: input.Code,
		Name: input.Name,
	}
	_, exist := mapSuppliers[id]
	if !exist {
		return nil, errors.New("supplier not found")
	}
	mapSuppliers[id].Code = input.Code
	mapSuppliers[id].Name = input.Name
	return supplier, nil
}

func (r *queryResolver) Suppliers(ctx context.Context) ([]*Supplier, error) {
	var suppliers []*Supplier
	for k := range mapSuppliers {
		suppliers = append(suppliers, mapSuppliers[k])
	}
	return suppliers, nil
}
