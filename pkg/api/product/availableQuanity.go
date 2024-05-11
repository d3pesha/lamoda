package product

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/api/errors"
	"net/http"
	"strconv"
)

// @Summary		Get Available Quantity
// @Accept		json
// @Produce		json
// @Tags		Product
// @Param		warehouseID	path	int	true	"warehouse ID"
// @Success	200 {array}  	model.ProductWarehouse "List of products"
// @Failure	400	{product} 	errors.APIError
// @Failure	500	{product}	errors.APIError
// @Router		/product/available/{id} [get]
func (r ProductRoute) GetAvailableProducts(c *gin.Context) {
	warehouseID := c.Param("id")

	id, err := strconv.Atoi(warehouseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid warehouse ID"})
		return
	}

	quantity, err := r.uc.GetAvailableProducts(c, uint32(id))
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": quantity})
}
