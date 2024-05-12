package product

import (
	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
	"lamoda/pkg/api/errors"
	"net/http"
)

// @Summary		List Products
// @Accept		json
// @Produce		json
// @Tags		Product
// @Success	200	{array}  	model.Product "List of products"
// @Failure	500	{product} 	errors.APIError
// @Router		/product/ [get]
func (r ProductRoute) ListProducts(c *gin.Context) {
	response, err := r.uc.GetAll(c.Request.Context())
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
