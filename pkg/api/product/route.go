package product

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/service"
)

type ProductRoute struct {
	uc service.ProductUseCase
}

func NewProductRoute(uc service.ProductUseCase) *ProductRoute {
	return &ProductRoute{
		uc: uc,
	}
}

func (r ProductRoute) Register(router *gin.Engine) {
	product := router.Group("/product")
	{
		product.POST("/create", r.Create)
		product.POST("/reserve", r.ReserveProduct)
		product.POST("/release", r.ReleaseReservedProducts)

		product.GET("/:id", r.GetById)
		product.GET("/", r.ListClients)
		product.GET("/available/:id", r.GetAvailableProducts)
	}

}
