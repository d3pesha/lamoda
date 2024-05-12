package warehouse

import (
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
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
	warehouse := &model.WarehouseCreateReq{}

	if err := jsonapi.UnmarshalPayload(c.Request.Body, warehouse); err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": aerr})
		return
	}

	response, err := r.uc.Create(c.Request.Context(), warehouse)
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}
	c.Writer.Header().Set("Content-Type", jsonapi.MediaType)
	c.Writer.WriteHeader(http.StatusCreated)

	if err = jsonapi.MarshalPayloadWithoutIncluded(c.Writer, response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error marshalling JSON API response"})
	}
}
