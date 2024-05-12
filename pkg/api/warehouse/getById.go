package warehouse

import (
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"lamoda/pkg/api/errors"
	"net/http"
	"strconv"
)

// @Summary		Get Warehouse by ID
// @Accept		json
// @Produce		json
// @Tags		Warehouse
// @Param		id	path	int	true	"Warehouse ID"
// @Success	200	{warehouse} 	model.Warehouse "Warehouse details"
// @Failure	400	{warehouse} 	errors.APIError
// @Failure	500	{warehouse} 	errors.APIError
// @Router		/warehouse/{id} [get]
func (r WarehouseRoute) GetWarehouseByID(c *gin.Context) {
	warehouseID := c.Param("id")

	id, err := strconv.Atoi(warehouseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uc ID"})
		return
	}

	response, err := r.uc.GetByID(c.Request.Context(), uint32(id))
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.Writer.Header().Set("Content-Type", jsonapi.MediaType)
	c.Writer.WriteHeader(http.StatusOK)

	if err = jsonapi.MarshalPayload(c.Writer, response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshalling JSON API response"})
	}
}
