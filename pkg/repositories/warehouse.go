package repositories

import (
	"context"
	log "github.com/sirupsen/logrus"
	"lamoda/pkg/data"
	"lamoda/pkg/model"
)

type warehouseRepo struct {
	data *data.Data
	log  *log.Logger
}

func NewStorageRepo(data *data.Data, logger *log.Logger) WarehouseRepo {
	return &warehouseRepo{
		data: data,
		log:  logger,
	}
}

type Warehouse struct {
	Id        uint32    `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Available *bool     `gorm:"column:available"`
	Products  []Product `gorm:"many2many:product_warehouses;"`
}

func (Warehouse) TableName() string {
	return "warehouses"
}

func (w Warehouse) modelToResponse() *model.Warehouse {
	dto := &model.Warehouse{
		Id:        w.Id,
		Name:      w.Name,
		Available: w.Available,
	}

	if w.Products != nil {
		for _, product := range w.Products {
			dto.Products = append(dto.Products, product.modelToResponse())
		}

	}
	return dto
}

func (r warehouseRepo) Create(_ context.Context, storage *model.WarehouseCreateReq) (*model.Warehouse, error) {
	var storageInfo Warehouse

	storageInfo.Name = storage.Name
	storageInfo.Available = storage.Available

	result := r.data.Db.Create(&storageInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return storageInfo.modelToResponse(), nil
}

func (r warehouseRepo) Update(_ context.Context, warehouse *model.Warehouse) (*model.Warehouse, error) {
	var warehouseInfo Warehouse

	warehouseInfo.Name = warehouse.Name
	warehouseInfo.Available = warehouse.Available

	result := r.data.Db.Updates(&warehouseInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return warehouseInfo.modelToResponse(), nil
}

func (r warehouseRepo) Delete(_ context.Context, id uint32) error {
	return r.data.Db.Model(&Warehouse{}).Delete(id).Error
}

func (r warehouseRepo) GetAll(_ context.Context) ([]*model.Warehouse, error) {
	var storages []*Warehouse

	result := r.data.Db.Find(&storages)
	if result.Error != nil {
		return nil, result.Error
	}

	storageResponse := make([]*model.Warehouse, 0)
	for _, user := range storages {
		storageResponse = append(storageResponse, user.modelToResponse())
	}

	return storageResponse, nil
}

func (r warehouseRepo) GetByID(_ context.Context, id uint32) (*model.Warehouse, error) {
	var storage Warehouse

	result := r.data.Db.First(&storage, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return storage.modelToResponse(), nil
}
