package repositories

import (
	"context"
	"lamoda/pkg/data"
	"lamoda/pkg/model"

	log "github.com/sirupsen/logrus"
)

type productRepo struct {
	data *data.Data
	log  *log.Logger
}

func NewProductRepo(data *data.Data, logger *log.Logger) ProductRepo {
	return &productRepo{
		data: data,
		log:  logger,
	}
}

type Product struct {
	ID         uint32              `gorm:"column:id"`
	Name       string              `gorm:"column:name"`
	Code       string              `gorm:"column:code"`
	Quantity   uint32              `gorm:"column:quantity"`
	Size       uint16              `gorm:"column:size"`
	Warehouses []*ProductWarehouse `gorm:"many2many:product_warehouses;"`
}

type ProductWarehouse struct {
	ID          uint32 `gorm:"column:id"`
	WarehouseID uint32 `gorm:"column:warehouse_id"`
	ProductID   uint32 `gorm:"column:product_id"`
	Quantity    uint32 `gorm:"column:quantity"`
}

type Reservation struct {
	ID          uint32 `gorm:"column:id"`
	WarehouseID uint32 `gorm:"column:warehouse_id"`
	ProductID   uint32 `gorm:"column:product_id"`
	Quantity    uint32 `gorm:"column:quantity"`
}

func (Product) TableName() string {
	return "products"
}

func (ProductWarehouse) TableName() string {
	return "product_warehouses"
}

func (Reservation) TableName() string {
	return "reservation"
}

func (p Product) modelToResponse() *model.Product {
	dto := &model.Product{
		ID:       p.ID,
		Name:     p.Name,
		Code:     p.Code,
		Quantity: p.Quantity,
		Size:     p.Size,
	}

	return dto
}

func (r ProductWarehouse) modelToResponse() *model.ProductWarehouse {
	dto := &model.ProductWarehouse{
		ID:          r.ID,
		WarehouseID: r.WarehouseID,
		ProductID:   r.ProductID,
		Quantity:    r.Quantity,
	}

	return dto
}

func (r Reservation) modelToResponse() *model.Reservation {
	dto := &model.Reservation{
		ID:          r.ID,
		WarehouseID: r.WarehouseID,
		ProductID:   r.ProductID,
		Quantity:    r.Quantity,
	}

	return dto
}

// Создание продукта, если были переданы данные по складам, то создается запись в таблице product_warehouses
func (r productRepo) Create(_ context.Context, product *model.ProductCreateReq) (*model.Product, error) {
	tx := r.data.Db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	productInfo := &Product{
		Name:     product.Name,
		Code:     product.Code,
		Quantity: product.Quantity,
		Size:     product.Size,
	}

	result := tx.Model(&Product{}).Create(&productInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	if product.Warehouses != nil && len(product.Warehouses) > 0 {
		warehouses := make([]*ProductWarehouse, 0)

		for _, w := range product.Warehouses {
			warehouse := ProductWarehouse{
				WarehouseID: w.ID,
				ProductID:   productInfo.ID,
				Quantity:    productInfo.Quantity,
			}

			warehouses = append(warehouses, &warehouse)
			productInfo.Warehouses = append(productInfo.Warehouses, &warehouse)
		}

		result = tx.Model(&ProductWarehouse{}).Create(&warehouses)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return productInfo.modelToResponse(), nil
}

// получаем список существующих продуктов
func (r productRepo) GetAll(_ context.Context) ([]*model.Product, error) {
	var products []Product

	result := r.data.Db.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	productResponse := make([]*model.Product, 0)

	for _, product := range products {
		productResponse = append(productResponse, product.modelToResponse())
	}

	return productResponse, nil
}

// получение продукта по id
func (r productRepo) GetByID(_ context.Context, id uint32) (*model.Product, error) {
	var product Product

	result := r.data.Db.Where("id = ?", id).First(&product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product.modelToResponse(), nil
}

// получаем доступные продукты на заданном складе, условия != 0 и склад доступен
func (r productRepo) GetAvailableProducts(_ context.Context, warehouseID uint32) ([]*model.ProductWarehouse, error) {
	var products []ProductWarehouse

	result := r.data.Db.Table("product_warehouses").
		Joins("JOIN warehouses ON product_warehouses.warehouse_id = warehouses.id").
		Where("product_warehouses.warehouse_id = ? AND product_warehouses.quantity > 0 AND warehouses.available = true", warehouseID).
		Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	response := make([]*model.ProductWarehouse, len(products))

	for i, product := range products {
		response[i] = product.modelToResponse()
	}

	return response, nil
}

// Резерв товара
func (r productRepo) Reserve(_ context.Context, reserve []*model.Reservation) error {
	tx := r.data.Db.Model(Reservation{}).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, res := range reserve {
		reserveInfo := &Reservation{
			ID:          res.ID,
			WarehouseID: res.WarehouseID,
			ProductID:   res.ProductID,
			Quantity:    res.Quantity,
		}

		// Если запись о резерве есть, то обновляем количество зарезервированного товара
		if reserveInfo.ID != 0 {
			if err := tx.Where("id = ?", reserveInfo.ID).
				Update("quantity", &reserveInfo.Quantity).Error; err != nil {
				tx.Rollback()
				return err
			}
			//в ином случае создаем новую запись
		} else {
			if err := tx.Where("product_id = ? AND warehouse_id = ?", reserveInfo.ProductID, reserveInfo.WarehouseID).
				Create(&reserveInfo).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r productRepo) ReleaseReserve(_ context.Context, reserve []*model.Reservation) error {
	tx := r.data.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, res := range reserve {
		if err := tx.Model(&Reservation{}).Where("product_id = ? AND warehouse_id = ?", res.ProductID, res.WarehouseID).
			Update("quantity", res.Quantity).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r productRepo) GetProductAndWarehouse(_ context.Context, productID []uint32, warehouseID []uint32) ([]*model.ProductWarehouse, error) {
	var productWarehouse []*ProductWarehouse

	for i, id := range productID {
		var product ProductWarehouse

		result := r.data.Db.Model(&ProductWarehouse{}).Where("product_id = ? AND warehouse_id = ?", id, warehouseID[i]).
			First(&product)
		if result.Error != nil {
			return nil, result.Error
		}

		productWarehouse = append(productWarehouse, &product)
	}

	response := make([]*model.ProductWarehouse, len(productWarehouse))
	for i, pw := range productWarehouse {
		response[i] = pw.modelToResponse()
	}

	return response, nil
}

func (r productRepo) UpdateProductWarehouse(_ context.Context, warehouse model.ProductWarehouse) error {
	tx := r.data.Db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Model(&ProductWarehouse{}).Where("product_id = ? AND warehouse_id = ?", warehouse.ProductID, warehouse.WarehouseID).
		Update("quantity", warehouse.Quantity).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r productRepo) GetReserve(_ context.Context, productID uint32, warehouseID uint32) (*model.Reservation, error) {
	var reserve Reservation

	err := r.data.Db.Model(&Reservation{}).Where("product_id = ? AND warehouse_id = ?", productID, warehouseID).First(&reserve).Error
	if err != nil {
		return nil, err
	}

	return reserve.modelToResponse(), nil
}

func (r productRepo) DeleteReserve(_ context.Context, id uint32) error {
	return r.data.Db.Delete(&Reservation{}, id).Error
}
