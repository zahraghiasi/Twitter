package access

import (
	"github.com/ghiac/social/pkg/text"
	"net/http"
)

type ServerError struct {
	Response    text.ServerText
	Description *string
	Code        int
	Where       string
	Error       error
}

func GetInternalError(where string, err error) *ServerError {
	return &ServerError{
		Response: text.TryAgain,
		Code:     http.StatusInternalServerError,
		Where:    where,
		Error:    err,
	}
}

func GetError(code text.ServerText) *ServerError {
	switch code {
	case text.NotEnoughChar:
		return &ServerError{
			Response: text.NotEnoughChar,
			Code:     http.StatusBadRequest,
		}
	case text.NotReceivedFile:
		return &ServerError{
			Response: text.NotReceivedFile,
			Code:     http.StatusBadRequest,
		}
	case text.InputDuplicated:
		return &ServerError{
			Response: text.NotReceivedFile,
			Code:     http.StatusBadRequest,
		}
	case text.UserOrPassInvalid:
		return &ServerError{
			Response: text.UserOrPassInvalid,
			Code:     http.StatusBadRequest,
		}
	case text.UserNotFound:
		return &ServerError{
			Response: text.UserNotFound,
			Code:     http.StatusNotFound,
		}
	case text.NotAcceptable:
		return &ServerError{
			Response: text.NotAcceptable,
			Code:     http.StatusNotAcceptable,
		}
	case text.NotFound:
		return &ServerError{
			Response: text.NotFound,
			Code:     http.StatusNotFound,
		}
	case text.InvalidInput:
		return &ServerError{
			Response: text.InvalidInput,
			Code:     http.StatusBadRequest,
		}
	case text.BadRequest:
		return &ServerError{
			Response: text.BadRequest,
			Code:     http.StatusBadRequest,
		}
	case text.Forbidden:
		return &ServerError{
			Response: text.Forbidden,
			Code:     http.StatusForbidden,
		}
	case text.NotDefined:
		return &ServerError{
			Response: text.NotDefined,
			Code:     http.StatusNotAcceptable,
		}
	}
	return &ServerError{
		Response: text.NotDefined,
		Code:     http.StatusServiceUnavailable,
	}
}
