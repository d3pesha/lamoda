package warehouse

import (
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"lamoda/pkg/api/errors"
	"net/http"
)

func (r WarehouseRoute) ListWarehouse(c *gin.Context) {
	response, err := r.uc.GetAll(c)
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
