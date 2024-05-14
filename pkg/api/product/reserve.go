package product

import (
	"lamoda/pkg/api/errors"
	"lamoda/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/jsonapi"
)

func (r ProductRoute) ReserveProduct(c *gin.Context) {
	reservations := &model.ReservationReq{}

	if err := jsonapi.UnmarshalPayload(c.Request.Body, reservations); err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": aerr})
		return
	}

	err := r.uc.Reserve(c.Request.Context(), reservations.Data)
	if err != nil {
		aerr := errors.DefaultErrorDecoder(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": aerr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
