package repositories

import (
	"context"
	"lamoda/pkg/model"
)

type WarehouseRepo interface {
	Create(context.Context, *model.WarehouseCreateReq) (*model.Warehouse, error)
	GetAll(context.Context) ([]*model.Warehouse, error)
	GetByID(context.Context, uint32) (*model.Warehouse, error)
}

type ProductRepo interface {
	Create(context.Context, *model.ProductCreateReq) (*model.Product, error)
	GetAll(context.Context) ([]*model.Product, error)
	GetByID(context.Context, uint32) (*model.Product, error)
	GetAvailableProducts(context.Context, uint32) ([]*model.ProductWarehouse, error)

	Reserve(context.Context, []*model.Reservation) error
	ReleaseReserve(context.Context, []*model.Reservation) error
	DeleteReserve(context.Context, uint32) error
	UpdateProductWarehouse(context.Context, model.ProductWarehouse) error
	GetProductAndWarehouse(context.Context, []uint32, []uint32) ([]*model.ProductWarehouse, error)
	GetReserve(context.Context, uint32, uint32) (*model.Reservation, error)
}
