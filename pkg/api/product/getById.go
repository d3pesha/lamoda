package product

import (
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"lamoda/pkg/api/errors"
	"net/http"
	"strconv"
)

func (r ProductRoute) GetById(c *gin.Context) {
	productID := c.Param("id")

	id, err := strconv.Atoi(productID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product ID"})
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
