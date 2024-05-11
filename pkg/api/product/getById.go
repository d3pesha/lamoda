package product

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/api/errors"
	"net/http"
	"strconv"
)

// @Summary		Get Product by ID
// @Accept		json
// @Produce		json
// @Tags		Product
// @Param		id	path	int	true		"Product ID"
// @Success	200	{product}	model.Product	"Product details"
// @Failure	400	{product} 	errors.APIError
// @Failure	500	{product} 	errors.APIError
// @Router		/product/{id} [get]
func (r ProductRoute) GetById(c *gin.Context) {
	productID := c.Param("id")
	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
		return
	}

	product, err := r.uc.GetByID(c.Request.Context(), uint32(id))
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.JSON(http.StatusOK, product)
}
