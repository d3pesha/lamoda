package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	log "github.com/sirupsen/logrus"
	"lamoda/pkg/model"
	"lamoda/pkg/repositories"
)

type ProductUseCase struct {
	repo   repositories.ProductRepo
	logger *log.Logger
}

func NewProductUseCase(repo repositories.ProductRepo, logger *log.Logger) *ProductUseCase {
	return &ProductUseCase{logger: logger, repo: repo}
}

var (
	ProductCreateError = errors.InternalServer("PRODUCT_CREATE_ERROR", "product create error")
	ProductUpdateError = errors.NotFound("PRODUCT_UPDATE_ERROR", "product update error")
	ProductNotFound    = errors.NotFound("PRODUCT_NOT_FOUND", "product not found")
	ProductFoundError  = errors.InternalServer("PRODUCT_FOUND_ERROR", "product found error")

	ReserveQuantityError = errors.InternalServer("RESERVE_QUANTITY_ERROR", "insufficient products for reserve")
	ReserveFoundError    = errors.InternalServer("RESERVE_FOUND_ERROR", "reserve found error")
	ReserveCreateError   = errors.InternalServer("RESERVE_CREATE_ERROR", "reserve create error")
	ReserveReleaseError  = errors.InternalServer("RESERVE_RELEASE_ERROR", "reserve release error")
	ReserveDeleteError   = errors.InternalServer("RESERVE_DELETE_ERROR", "reserve delete error")
)

func (uc ProductUseCase) Create(ctx context.Context, product *model.ProductCreateReq) (*model.Product, error) {
	response, err := uc.repo.Create(ctx, product)
	if err != nil {
		return nil, ProductCreateError
	}

	return response, nil
}

func (uc ProductUseCase) GetAll(ctx context.Context) ([]*model.Product, error) {
	return uc.repo.GetAll(ctx)
}

func (uc ProductUseCase) GetByID(ctx context.Context, id uint32) (*model.Product, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc ProductUseCase) GetAvailableProducts(ctx context.Context, warehouseID uint32) ([]*model.ProductWarehouse, error) {
	products, err := uc.repo.GetAvailableQuantity(ctx, warehouseID)
	if err != nil {
		return nil, ProductFoundError
	}

	return products, nil
}

func (uc ProductUseCase) Reserve(ctx context.Context, reserve []*model.Reservation) error {
	productIds := make([]uint32, len(reserve))
	warehouseIds := make([]uint32, len(reserve))
	reserveQuantity := make([]uint32, len(reserve))

	for i, r := range reserve {
		productIds[i] = r.ProductID
		reserveQuantity[i] = r.Quantity
		warehouseIds[i] = r.WarehouseID
	}

	// получаем список продуктов на складе
	products, err := uc.repo.GetProductAndWarehouse(ctx, productIds, warehouseIds)
	if err != nil {
		return ProductFoundError
	}

	reservations := make([]*model.Reservation, len(reserve))

	for i, product := range products {
		// Проверяем чтобы количество запрошенного товара не превышало остаток на складе
		if product.Quantity < reserveQuantity[i] {
			return ReserveQuantityError
		}

		updateProduct := model.ProductWarehouse{
			ProductID:   product.ProductID,
			Quantity:    product.Quantity - reserveQuantity[i],
			WarehouseID: product.WarehouseID,
		}

		// проверяем наличие записи резерва
		existReserve, _ := uc.repo.GetReserve(ctx, product.ProductID, product.WarehouseID)

		if existReserve != nil {
			// если запись есть, то обновляем данные
			reserveQuantity[i] += existReserve.Quantity
		}

		// вызываем обновление количества товара на складе
		err = uc.repo.UpdateProductWarehouse(ctx, updateProduct)
		if err != nil {
			return ProductUpdateError
		}

		newReservation := &model.Reservation{
			ProductID:   product.ProductID,
			Quantity:    reserveQuantity[i],
			WarehouseID: product.WarehouseID,
		}

		// ставим id резерва для обновления существующего резерва
		if existReserve != nil {
			newReservation.ID = existReserve.ID
		}

		reservations[i] = newReservation
	}

	err = uc.repo.Reserve(ctx, reservations)
	if err != nil {
		return ReserveCreateError
	}

	return nil
}

func (uc ProductUseCase) ReleaseReserve(ctx context.Context, reserve []*model.Reservation) error {
	reservations := make([]*model.Reservation, 0)

	for _, r := range reserve {

		// проверка на существующий резерв
		existReserve, err := uc.repo.GetReserve(ctx, r.ProductID, r.WarehouseID)
		if err != nil {
			return ReserveFoundError
		}

		// если количество зарезервированного товара меньше, чем запрошенного - возвращаем ошибку
		if existReserve.Quantity < r.Quantity {
			return ReserveQuantityError
		}

		// если количество товара в резерве совпадает с количеством заявленного для выпуска, то удаляем запись о резерве
		if existReserve.Quantity == r.Quantity {
			err = uc.repo.DeleteReserve(ctx, existReserve.ID)
			if err != nil {
				return ReserveDeleteError
			}

			// в ином случае обновляем количество зарезервированного товара на определенном складе
		} else {
			updateReservation := &model.Reservation{
				WarehouseID: r.WarehouseID,
				ProductID:   r.ProductID,
				Quantity:    existReserve.Quantity - r.Quantity,
			}
			reservations = append(reservations, updateReservation)
		}
	}

	err := uc.repo.ReleaseReserve(ctx, reservations)
	if err != nil {
		return ReserveReleaseError
	}

	return nil
}

func (uc ProductUseCase) DeleteReserve(ctx context.Context, id uint32) error {
	err := uc.repo.DeleteReserve(ctx, id)
	if err != nil {
		return ReserveDeleteError
	}

	return nil
}

func (uc ProductUseCase) UpdateProductAndWarehouse(ctx context.Context, warehouse model.ProductWarehouse) error {
	return uc.repo.UpdateProductWarehouse(ctx, warehouse)
}

func (uc ProductUseCase) GetProductAndWarehouse(ctx context.Context, productID []uint32, storageID []uint32) ([]*model.ProductWarehouse, error) {
	productWarehouses, err := uc.repo.GetProductAndWarehouse(ctx, productID, storageID)
	if err != nil {
		return nil, ProductFoundError
	}

	return productWarehouses, nil
}

func (uc ProductUseCase) GetReserve(ctx context.Context, productID uint32, warehouseID uint32) (*model.Reservation, error) {
	return uc.repo.GetReserve(ctx, productID, warehouseID)
}
