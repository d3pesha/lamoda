package warehouse

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/api/errors"
	"net/http"
)

// @Summary		List Warehouses
// @Accept		json
// @Produce		json
// @Tags		Warehouse
// @Success	200	{array} model.Warehouse "List of warehouse"
// @Failure	500	{warehouse} 	errors.APIError
// @Router		/warehouse/ [get]
func (r WarehouseRoute) ListWarehouse(c *gin.Context) {
	storages, err := r.uc.GetAll(c)
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.JSON(http.StatusOK, storages)
}
