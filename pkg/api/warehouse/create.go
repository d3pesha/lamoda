package warehouse

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/api/errors"
	"lamoda/pkg/model"
	"net/http"
)

// @Summary		Create Warehouse
// @Accept		json
// @Produce		json
// @Tags		Warehouse
// @Param		warehouse	body	model.WarehouseCreateReq	true	"Warehouse"
// @Success	201
// @Failure	400	{warehouse} 	errors.APIError
// @Failure	500	{warehouse} 	errors.APIError
// @Router		/warehouse/ [post]
func (r WarehouseRoute) Create(c *gin.Context) {
	var warehouse model.WarehouseCreateReq

	if err := c.BindJSON(&warehouse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := r.uc.Create(c, &warehouse)
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"warehouse": response})
}
