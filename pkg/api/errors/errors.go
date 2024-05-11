package errors

import (
	"github.com/go-kratos/kratos/v2/errors"
)

type APIError struct {
	Code    int    `json:"code"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

func DefaultErrorDecoder(err error) *APIError {
	e := errors.FromError(err)
	if e == nil {
		return &APIError{
			Code:    500,
			Key:     errors.UnknownReason,
			Message: errors.UnknownReason,
		}
	} else {
		return &APIError{
			Code:    int(e.Code),
			Key:     e.Reason,
			Message: e.Message,
		}
	}
}
