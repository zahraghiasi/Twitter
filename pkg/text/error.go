package text

func GetError(cons ServerText, lang ServerLang) string {
	var errors = map[ServerText]WordLang{
		ExceedLimit: {
			English: "Exceed Limit",
			Persian: "محدودیت شما به اتمام رسیده است",
		},
		BadRequest: {
			English: "invalid input",
			Persian: "ورودی نامعتبر است",
		},
		NotFound: {
			English: "پیدا نشد",
			Persian: "Not found",
		},
		TryAgain: {
			English: "Try Again",
			Persian: "دوباره تلاش کنید",
		},
		InvalidInput: {
			English: "Invalid Input",
			Persian: "ورودی نامعتبر است",
		},
		UserNotFound: {
			English: "UserNotFound",
			Persian: "کاربر پیدا نشد",
		},
		NotEnoughChar: {
			English: "NotEnoughChar",
			Persian: "تعداد کارکتر از حد مجاز کمتر است",
		},
		NotReceivedFile: {
			English: "NotReceivedFile",
			Persian: "فایل آپلود نشده",
		},
		NotAcceptable: {
			English: "NotAcceptable",
			Persian: "غیرقابل قبول",
		},
		Forbidden: {
			English: "Forbidden",
			Persian: "ورود غیرمجاز است",
		},
		TitleAddressInvalid: {
			English: "Invalid title",
			Persian: "عنوان مکان کمتر از حد مجاز است",
		},
		UsernameExist: {
			English: "Username Exist!",
			Persian: "نام کاربری وجود دارد",
		},
		UserOrPassInvalid: {
			English: "UserOrPassInvalid!",
			Persian: "UserOrPassInvalid",
		},
	}
	if lang == Fa {
		return errors[cons].Persian
	} else {
		return errors[cons].English
	}
}

type ServerText int

type WordLang struct {
	English string
	Persian string
}

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
