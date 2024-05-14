package warehouse

import (
	"lamoda/pkg/api/errors"
	"lamoda/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
)

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
