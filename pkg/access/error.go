package access

import (
	"github.com/ghiac/social/pkg/text"
	"net/http"
)

type ServerError struct {
	Response    text.ServerText
	Description *text.WordLang
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
	case text.ProductNotFound:
		return &ServerError{
			Response: text.ProductNotFound,
			Code:     http.StatusNotFound,
		}
	case text.CantIncreaseShouldNewOrder:
		return &ServerError{
			Response: text.CantIncreaseShouldNewOrder,
			Code:     http.StatusBadRequest,
		}
	case text.CantChangePaidOrders:
		return &ServerError{
			Response: text.CantChangePaidOrders,
			Code:     http.StatusNotAcceptable,
		}
	case text.CostumerNotSelected:
		return &ServerError{
			Response: text.CostumerNotSelected,
			Code:     http.StatusBadRequest,
		}
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
	case text.AddressNotFound:
		return &ServerError{
			Response: text.AddressNotFound,
			Code:     http.StatusNotFound,
		}
	case text.ExistDiscountInTime:
		return &ServerError{
			Response: text.ExistDiscountInTime,
			Code:     http.StatusNotAcceptable,
		}
	case text.TableNotFound:
		return &ServerError{
			Response: text.TableNotFound,
			Code:     http.StatusForbidden,
		}
	case text.CantChangeSubProduct:
		return &ServerError{
			Response: text.CantChangeSubProduct,
			Code:     http.StatusBadRequest,
		}
	case text.DeliveryDisabled:
		return &ServerError{
			Response: text.DeliveryDisabled,
			Code:     http.StatusNotAcceptable,
		}
	case text.InvalidLocation:
		return &ServerError{
			Response: text.CashDeskIsOpen,
			Code:     http.StatusBadRequest,
		}
	case text.CashDeskIsOpen:
		return &ServerError{
			Response: text.CashDeskIsOpen,
			Code:     http.StatusNotAcceptable,
		}
	case text.OutOfServiceArea:
		return &ServerError{
			Response: text.OutOfServiceArea,
			Code:     http.StatusNotAcceptable,
		}
	case text.InputTimeInvalid:
		return &ServerError{
			Response: text.InputTimeInvalid,
			Code:     400,
		}
	case text.UserNotFound:
		return &ServerError{
			Response: text.UserNotFound,
			Code:     http.StatusNotFound,
		}
	case text.TableHasNotPaidOrders:
		return &ServerError{
			Response: text.TableHasNotPaidOrders,
			Code:     http.StatusNotAcceptable,
		}
	case text.OrderNotClosed:
		return &ServerError{
			Response: text.OrderNotClosed,
			Code:     http.StatusNotAcceptable,
		}
	case text.SessionSelected:
		return &ServerError{
			Response: text.SessionSelected,
			Code:     http.StatusNotAcceptable,
		}
	case text.SessionDisabled:
		return &ServerError{
			Response: text.SessionDisabled,
			Code:     http.StatusNotAcceptable,
		}
	case text.SessionInPayment:
		return &ServerError{
			Response: text.SessionInPayment,
			Code:     http.StatusNotAcceptable,
		}
	case text.SessionReserved:
		return &ServerError{
			Response: text.SessionReserved,
			Code:     http.StatusNotAcceptable,
		}
	case text.TimePast:
		return &ServerError{
			Response: text.TimePast,
			Code:     http.StatusNotAcceptable,
		}
	case text.SessionNotFound:
		return &ServerError{
			Response: text.SessionNotFound,
			Code:     http.StatusNotFound,
		}
	case text.SumValuesNotEqual:
		return &ServerError{
			Response: text.SumValuesNotEqual,
			Code:     http.StatusBadRequest,
		}
	case text.OrderNotFound:
		return &ServerError{
			Response: text.OrderNotFound,
			Code:     http.StatusNotFound,
		}
	case text.NotAcceptable:
		return &ServerError{
			Response: text.NotAcceptable,
			Code:     http.StatusNotAcceptable,
		}
	case text.TitleCharIsNotEnough:
		return &ServerError{
			Response: text.TitleCharIsNotEnough,
			Code:     http.StatusBadRequest,
		}
	case text.NotFound:
		return &ServerError{
			Response: text.NotFound,
			Code:     http.StatusNotFound,
		}
	case text.SelectJustOneSession:
		return &ServerError{
			Response: text.SelectJustOneSession,
			Code:     http.StatusBadRequest,
		}
	case text.SessionFilled:
		return &ServerError{
			Response: text.SessionFilled,
			Code:     http.StatusNotAcceptable,
		}
	case text.InvalidInput:
		return &ServerError{
			Response: text.InvalidInput,
			Code:     http.StatusBadRequest,
		}
	case text.RequestTypeNotFound:
		return &ServerError{
			Response: text.RequestTypeNotFound,
			Code:     http.StatusNotFound,
		}
	case text.ShopNotFound:
		return &ServerError{
			Response: text.ShopNotFound,
			Code:     http.StatusNotFound,
		}
	case text.ShopClosed:
		return &ServerError{
			Response: text.ShopClosed,
			Code:     http.StatusNotAcceptable,
		}
	case text.TableOrderNotFound:
		return &ServerError{
			Response: text.TableOrderNotFound,
			Code:     http.StatusNotFound,
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
	case text.TransactionExpired:
		return &ServerError{
			Response: text.TransactionExpired,
			Code:     http.StatusNotAcceptable,
		}
	case text.NotDefined:
		return &ServerError{
			Response: text.NotDefined,
			Code:     http.StatusNotAcceptable,
		}
	case text.TicketNotFound:
		return &ServerError{
			Response: text.TicketNotFound,
			Code:     http.StatusNotFound,
		}
	case text.TransactionFailed:
		return &ServerError{
			Response: text.TransactionFailed,
			Code:     http.StatusExpectationFailed,
		}
	case text.TransactionNotFound:
		return &ServerError{
			Response: text.TransactionNotFound,
			Code:     http.StatusNotFound,
		}
	}
	return &ServerError{
		Response: text.NotDefined,
		Code:     http.StatusServiceUnavailable,
	}
}
