package product

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/api/errors"
	"lamoda/pkg/model"
	"net/http"
)

// @Summary		Create
// @Accept		json
// @Produce		json
// @Tags		Product
// @Param		Product		body	model.Product	true	"Product"
// @Success	200 {product}  model.Product	"Product details"
// @Failure	500	{product}	errors.APIError
// @Failure	400	{product}	errors.APIError
// @Router		/product/ [post]
func (r ProductRoute) Create(c *gin.Context) {
	var product model.ProductCreateReq

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := r.uc.Create(c, &product)
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"product": response})
}
