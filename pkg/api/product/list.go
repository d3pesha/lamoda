package product

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/api/errors"
	"net/http"
)

// @Summary		List Products
// @Accept		json
// @Produce		json
// @Tags		Product
// @Success	200	{array}  	model.Product "List of products"
// @Failure	500	{product} 	errors.APIError
// @Router		/product/ [get]
func (r ProductRoute) ListClients(c *gin.Context) {
	products, err := r.uc.GetAll(c.Request.Context())
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.JSON(http.StatusOK, products)
}
