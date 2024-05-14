package service

import (
	"context"
	"lamoda/pkg/model"
	"lamoda/pkg/repositories"

	"github.com/go-kratos/kratos/v2/errors"
	log "github.com/sirupsen/logrus"
)

type WarehouseUseCase struct {
	repo   repositories.WarehouseRepo
	logger *log.Logger
}

func NewWarehouseUseCase(repo repositories.WarehouseRepo, logger *log.Logger) *WarehouseUseCase {
	return &WarehouseUseCase{logger: logger, repo: repo}
}

var (
	WarehouseCreateError = errors.InternalServer("WAREHOUSE_CREATE_ERROR", "warehouse create error")
	WarehouseNotFound    = errors.NotFound("WAREHOUSE_NOT_FOUND", "warehouse not found")
	WarehouseFoundError  = errors.InternalServer("WAREHOUSE_FOUND_ERROR", "warehouse found error")
)

func (uc WarehouseUseCase) Create(ctx context.Context, warehouse *model.WarehouseCreateReq) (*model.Warehouse, error) {
	response, err := uc.repo.Create(ctx, warehouse)
	if err != nil {
		return nil, WarehouseCreateError
	}

	return response, nil
}

func (uc WarehouseUseCase) GetAll(ctx context.Context) ([]*model.Warehouse, error) {
	warehouses, err := uc.repo.GetAll(ctx)
	if err != nil {
		return nil, WarehouseFoundError
	}

	return warehouses, nil
}

func (uc WarehouseUseCase) GetByID(ctx context.Context, id uint32) (*model.Warehouse, error) {
	warehouse, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, WarehouseNotFound
	}

	return warehouse, nil
}
