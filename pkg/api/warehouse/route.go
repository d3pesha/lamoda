package warehouse

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/service"
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
