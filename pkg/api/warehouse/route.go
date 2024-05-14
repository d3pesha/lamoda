package warehouse

import (
	"lamoda/pkg/service"

	"github.com/gin-gonic/gin"
)

type WarehouseRoute struct {
	uc service.WarehouseUseCase
}

func NewUserRoute(uc service.WarehouseUseCase) *WarehouseRoute {
	return &WarehouseRoute{
		uc: uc,
	}
}

func (r WarehouseRoute) Register(router *gin.Engine) {
	storage := router.Group("/warehouse")
	{
		storage.POST("/create", r.Create)
		storage.GET("/:id", r.GetWarehouseByID)
		storage.GET("/", r.ListWarehouse)
	}
}
