package product

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/api/errors"
	"lamoda/pkg/model"
	"net/http"
)

// @Summary		Release Reserved Products
// @Accept		json
// @Produce		json
// @Tags		Product
// @Param		ReleaseRequest	body	model.Reservation	true	"Release request"
// @Success	200
// @Failure	400	{product} 	errors.APIError
// @Failure	500	{product} 	errors.APIError
// @Router		/product/release [post]
func (r ProductRoute) ReleaseReservedProducts(c *gin.Context) {
	var request []*model.Reservation

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.uc.ReleaseReserve(c.Request.Context(), request)
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
