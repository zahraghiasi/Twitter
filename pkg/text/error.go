package text

func GetError(cons ServerText) string {
	var errors = map[ServerText]string{
		ExceedLimit:         "محدودیت شما به اتمام رسیده است",
		BadRequest:          "ورودی نامعتبر است",
		NotFound:            "Not found",
		TryAgain:            "دوباره تلاش کنید",
		InvalidInput:        "ورودی نامعتبر است",
		UserNotFound:        "کاربر پیدا نشد",
		NotEnoughChar:       "تعداد کارکتر از حد مجاز کمتر است",
		NotReceivedFile:     "فایل آپلود نشده",
		NotAcceptable:       "غیرقابل قبول",
		Forbidden:           "ورود غیرمجاز است",
		TitleAddressInvalid: "عنوان مکان کمتر از حد مجاز است",
		UsernameExist:       "نام کاربری وجود دارد",
		UserOrPassInvalid:   "UserOrPassInvalid",
	}
	return errors[cons]
}

type ServerText int

func (s ServerText) String() string {
	switch s {
	case NotFound:
		return "NotFound"
	case TryAgain:
		return "TryAgain"
	case InvalidInput:
		return "InvalidInput"
	case UserNotFound:
		return "UserNotFound"
	case NotEnoughChar:
		return "NotEnoughChar"
	case NotReceivedFile:
		return "NotReceivedFile"
	case NotAcceptable:
		return "NotAcceptable"
	case Forbidden:
		return "Forbidden"
	case BadRequest:
		return "BadRequest"
	case ExceedLimit:
		return "ExceedLimit"
	case TitleAddressInvalid:
		return "TitleAddressInvalid"
	case UsernameExist:
		return "UsernameExist"
	case InputDuplicated:
		return "InputDuplicated"
	default:
		return "NotDefined"
	}
}

func (s ServerText) FromString(e string) ServerText {
	switch e {
	case "NotFound":
		return NotFound
	case "TryAgain":
		return TryAgain
	case "InvalidInput":
		return InvalidInput
	case "UserNotFound":
		return UserNotFound
	case "NotEnoughChar":
		return NotEnoughChar
	case "NotReceivedFile":
		return NotReceivedFile
	case "NotAcceptable":
		return NotAcceptable
	case "Forbidden":
		return Forbidden
	case "BadRequest":
		return BadRequest
	case "ExceedLimit":
		return ExceedLimit
	case "TitleAddressInvalid":
		return TitleAddressInvalid
	case "UsernameExist":
		return UsernameExist
	case "UserOrPassInvalid":
		return UserOrPassInvalid
	case "InputDuplicated":
		return InputDuplicated
	}
	return NotDefined
}

const (
	NotFound ServerText = iota
	TryAgain
	InvalidInput
	UserNotFound
	NotEnoughChar
	NotReceivedFile
	NotAcceptable
	BadRequest
	ExceedLimit
	TitleAddressInvalid
	UsernameExist
	NotDefined
	Forbidden
	UserOrPassInvalid
	InputDuplicated
)
