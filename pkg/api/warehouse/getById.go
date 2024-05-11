package warehouse

import (
	"github.com/gin-gonic/gin"
	"lamoda/pkg/api/errors"
	"net/http"
	"strconv"
)

// @Summary		Get Warehouse by ID
// @Accept		json
// @Produce		json
// @Tags		Warehouses
// @Param		id	path	int	true	"Warehouse ID"
// @Success	200	{warehouse} model.Warehouse "Warehouse details"
// @Failure	400	{warehouse} 	errors.APIError
// @Failure	500	{warehouse} 	errors.APIError
// @Router		/warehouse/{id} [get]
func (r WarehouseRoute) GetWarehouseByID(c *gin.Context) {
	storageID := c.Param("id")

	id, err := strconv.Atoi(storageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uc ID"})
		return
	}

	warehouse, err := r.uc.GetByID(c.Request.Context(), uint32(id))
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.JSON(http.StatusOK, warehouse)
}
