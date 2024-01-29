package text

import "strconv"

func GetRangeError(s, e int, fa, en string, lang ServerLang) string {
	return "بازه ی قابل قبول برای " + fa + " " + strconv.Itoa(s) + " " + strconv.Itoa(e) + "می باشد."
}
func GetSmallError(fa, en string, lang ServerLang) string {
	return fa + " کوتاه تر از حد مجاز است"
}

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
		InputTimeInvalid: {
			English: "Invalid Input Time",
			Persian: "زمان ورودی نامعتبر است",
		},
		InvalidInput: {
			English: "Invalid Input",
			Persian: "ورودی نامعتبر است",
		},
		IsExistWorkTime: {
			English: "IsExistWorkTime",
			Persian: "زمان مورد نظر قبلا پر شده است",
		},
		UserNotFound: {
			English: "UserNotFound",
			Persian: "کاربر پیدا نشد",
		},
		AddressNotFound: {
			English: "AddressNotFound",
			Persian: "آدرس پیدا نشد",
		},
		WorkTimeNotFound: {
			English: "WorkTimeNotFound",
			Persian: "زمان پیدا نشد",
		},
		WorkTimeOutOfNumber: {
			English: "WorkTimeOutOfNumber",
			Persian: "زمان بیشتر از مقدار تعیین شده است",
		},
		WorkTimeFilled: {
			English: "WorkTimeFilled",
			Persian: "زمان کاری پر است",
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
		DateIsNotFree: {
			English: "DateIsNotFree",
			Persian: "در این بازه جلسه تایید شده ی دیگری وجود دارد",
		},
		UserOrPassInvalid: {
			English: "Username or password is not valid",
			Persian: "پسورد یا نام کاربری نامعتبر است",
		},
		Forbidden: {
			English: "Forbidden",
			Persian: "ورود غیرمجاز است",
		},
		InputBirthdayInvalid: {
			English: "InputBirthdayInvalid",
			Persian: "تاریخ تولد اشتباه است",
		},
		InputNationalCodeInvalid: {
			English: "InputNationalCodeInvalid",
			Persian: "کدملی نامعتبر است",
		},
		TitleAddressInvalid: {
			English: "Invalid title",
			Persian: "عنوان مکان کمتر از حد مجاز است",
		},
		InvalidLocation: {
			English: "InvalidLocation",
			Persian: "مکان نامعتبر است",
		},
		InvalidTitleInput: {
			English: "Invalid title",
			Persian: "عنوان ورودی نامعتبر است",
		},
		ExistDiscountInTime: {
			English: "برای این بازه تخفیف وجود دارد",
			Persian: "برای این بازه تخفیف وجود دارد",
		},
		TitleCharIsNotEnough: {
			English: "Title character is not enough",
			Persian: "(بیش از ۳ کارکتر وارد کنید) تعداد کاراکترهای عنوان کافی نیست",
		},
		PriorityInvalid: {
			English: "Priority is Invalid",
			Persian: "الویت وارد شده نامعتبر است",
		},
		WebsiteInvalid: {
			English: "Website is Invalid",
			Persian: "آدرس وبسایت وارد شده نامعتبر است",
		},
		WebsiteDuplicated: {
			English: "WebsiteDuplicated",
			Persian: "آدرس وبسایت وارد شده تکراری است",
		},
		InputDuplicated: {
			English: "Input is duplicated",
			Persian: "اطلاعات وارد شده تکراری است",
		},
		SessionExist: {
			English: "SessionExist",
			Persian: "جلسه وجود دارد",
		},
		SessionSelected: {
			English: "SessionSelected",
			Persian: "جلسه انتخاب شده است",
		},
		SessionNotFound: {
			English: "SessionNotFound",
			Persian: "جلسه پیدا نشد",
		},
		SessionFilled: {
			English: "SessionFilled",
			Persian: "جلسه پر شده است",
		},
		SessionReserved: {
			English: "SessionReserved",
			Persian: "جلسه رزرو شده است",
		},
		OverflowDay: {
			English: "OverflowDay",
			Persian: "جلسات از روز مورد نظر خارج می شوند",
		},
		SessionNotAvailable: {
			English: "SessionNotAvailable",
			Persian: "جلس در دسترس نیست",
		},
		SessionFree: {
			English: "SessionFree",
			Persian: "جلس خالی است",
		},
		SelectJustOneSession: {
			English: "SelectJustOneSession",
			Persian: "فقط یک جلسه انتخاب کنید",
		},
		SessionInPayment: {
			English: "SessionInPayment",
			Persian: "جلسه در حال پرداخت است",
		},
		SessionDisabled: {
			English: "SessionDisabled",
			Persian: "جلسه غیرفعال است",
		},
		InvalidPhoneNumber: {
			English: "InvalidPhoneNumber",
			Persian: "شماره تماس نامعتبر",
		},
		TimePast: {
			English: "The time is past ",
			Persian: "زمان گذشته است",
		},
		TransactionNotFound: {
			English: "Transaction Not Found",
			Persian: "تراکنش پیدا نشد",
		},
		TransactionExpired: {
			English: "Transaction Expired",
			Persian: "تراکنش به اتمام رسیده است",
		},
		TransactionFailed: {
			English: "Transaction Failed",
			Persian: "تراکنش ناموفق",
		},
		OrderNotFound: {
			English: "Order Not Found",
			Persian: "سبد خرید پیدا نشد",
		},
		ShopClosed: {
			English: "Shop is closed",
			Persian: "فروشگاه تعطیل است",
		},
		DeliveryDisabled: {
			English: "Delivery disabled",
			Persian: "بیرون بر رستوران تعطیل است",
		},
		OutOfServiceArea: {
			English: "OutOfServiceArea",
			Persian: "خارج از محدوده ی ثبت شده فروشگاه است",
		},
		CarInParking: {
			English: "CarInParking",
			Persian: "ماشین درون پارکینگ است",
		},
		ProductNotFound: {
			English: "Product Not Found!",
			Persian: "محصول پیدا نشد",
		},
		ShopNotFound: {
			English: "Shop Not Found!",
			Persian: "فروشگاه پیدا نشد",
		},
		TableNotFound: {
			English: "Table Not Found!",
			Persian: "میز پیدا نشد",
		},
		TableReserved: {
			English: "Table Reserved!",
			Persian: "میز رزور است",
		},
		ShouldOnePayment: {
			English: "You should one payment!",
			Persian: "لزوم خرید و تخفیف بار اول پرداخت آنلاین می باشد",
		},
		TicketNotFound: {
			English: "TicketNotFound!",
			Persian: "TicketNotFound",
		},
		PaymentPaid: {
			English: "PaymentPaid!",
			Persian: "PaymentPaid",
		},
		OrderNotClosed: {
			English: "OrderNotClosed!",
			Persian: "سبد خرید بسته نشده",
		},
		MaxLowerThanMin: {
			English: "MaxLowerThanMin!",
			Persian: "MaxLowerThanMin",
		},
		TableHasNotPaidOrders: {
			English: "TableHasNotPaidOrders!",
			Persian: "میز سفارش پرداخت نشده دارد",
		},
		CantChangePaidOrders: {
			English: "CantChangePaidOrders!",
			Persian: "امکان تغییر سفارشات پرداخت شده وجود ندارد",
		},
		RequestTypeNotFound: {
			English: "RequestTypeNotFound!",
			Persian: "RequestTypeNotFound",
		},
		TableOrderNotFound: {
			English: "TableOrderNotFound!",
			Persian: "TableOrderNotFound",
		},
		CashDeskIsOpen: {
			English: "CashDeskIsOpen!",
			Persian: "صندوق تسویه نشده است",
		},
		OriginTableIsFree: {
			English: "OriginTableIsFree!",
			Persian: "میز مبدا آزاد است",
		},
		DestinationTableIsReserved: {
			English: "DestinationTableIsReserved!",
			Persian: "میز مقصد رزور شده است",
		},
		CantChangeSubProduct: {
			English: "CantChangeSubProduct!",
			Persian: "قابلیت تغییر زیر سفارش محصول به صورت مستقیم وجود ندارد",
		},
		UsernameExist: {
			English: "Username Exist!",
			Persian: "نام کاربری وجود دارد",
		},
		SumValuesNotEqual: {
			English: "The sum of the values with the amount of the table is not equal !",
			Persian: "مجموع مقادیر با مبلغ دریافتی میز برابر نیست",
		},
		CantIncreaseShouldNewOrder: {
			English: "CantIncreaseShouldNewOrder!",
			Persian: "برای اضافه کردن محصول سفارش جدید ثبت کنید",
		},
		CostumerNotSelected: {
			English: "Costumer not selected!",
			Persian: "مشتری انتخاب نشده است",
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
	case InputTimeInvalid:
		return "InputTimeInvalid"
	case InvalidInput:
		return "InvalidInput"
	case IsExistWorkTime:
		return "IsExistWorkTime"
	case UserNotFound:
		return "UserNotFound"
	case AddressNotFound:
		return "AddressNotFound"
	case WorkTimeNotFound:
		return "WorkTimeNotFound"
	case WorkTimeOutOfNumber:
		return "WorkTimeOutOfNumber"
	case WorkTimeFilled:
		return "WorkTimeFilled"
	case NotEnoughChar:
		return "NotEnoughChar"
	case NotReceivedFile:
		return "NotReceivedFile"
	case NotAcceptable:
		return "NotAcceptable"
	case DateIsNotFree:
		return "DateIsNotFree"
	case UserOrPassInvalid:
		return "UserOrPassInvalid"
	case Forbidden:
		return "Forbidden"
	case InputBirthdayInvalid:
		return "InputBirthdayInvalid"
	case InputNationalCodeInvalid:
		return "InputNationalCodeInvalid"
	case BadRequest:
		return "BadRequest"
	case ExceedLimit:
		return "ExceedLimit"
	case TitleAddressInvalid:
		return "TitleAddressInvalid"
	case InvalidLocation:
		return "InvalidLocation"
	case InvalidTitleInput:
		return "InvalidTitleInput"
	case ExistDiscountInTime:
		return "ExistDiscountInTime"
	case TitleCharIsNotEnough:
		return "TitleCharIsNotEnough"
	case PriorityInvalid:
		return "PriorityInvalid"
	case WebsiteInvalid:
		return "WebsiteInvalid"
	case WebsiteDuplicated:
		return "WebsiteDuplicated"
	case InputDuplicated:
		return "InputDuplicated"
	case SessionExist:
		return "SessionExist"
	case SessionSelected:
		return "SessionSelected"
	case SessionNotFound:
		return "SessionNotFound"
	case SessionFilled:
		return "SessionFilled"
	case SessionReserved:
		return "SessionReserved"
	case OverflowDay:
		return "OverflowDay"
	case SessionNotAvailable:
		return "SessionNotAvailable"
	case SessionFree:
		return "SessionFree"
	case SelectJustOneSession:
		return "SelectJustOneSession"
	case SessionInPayment:
		return "SessionInPayment"
	case SessionDisabled:
		return "SessionDisabled"
	case InvalidPhoneNumber:
		return "InvalidPhoneNumber"
	case TimePast:
		return "TimePast"
	case TransactionNotFound:
		return "TransactionNotFound"
	case TransactionExpired:
		return "TransactionExpired"
	case TransactionFailed:
		return "TransactionFailed"
	case OrderNotFound:
		return "OrderNotFound"
	case ShopClosed:
		return "ShopClosed"
	case DeliveryDisabled:
		return "DeliveryDisabled"
	case OutOfServiceArea:
		return "OutOfServiceArea"
	case CarInParking:
		return "CarInParking"
	case ProductNotFound:
		return "ProductNotFound"
	case ShopNotFound:
		return "ShopNotFound"
	case TableNotFound:
		return "TableNotFound"
	case TableReserved:
		return "TableReserved"
	case ShouldOnePayment:
		return "ShouldOnePayment"
	case TicketNotFound:
		return "TicketNotFound"
	case PaymentPaid:
		return "PaymentPaid"
	case OrderNotClosed:
		return "OrderNotClosed"
	case MaxLowerThanMin:
		return "MaxLowerThanMin"
	case TableHasNotPaidOrders:
		return "TableHasNotPaidOrders"
	case CantChangePaidOrders:
		return "CantChangePaidOrders"
	case RequestTypeNotFound:
		return "RequestTypeNotFound"
	case TableOrderNotFound:
		return "TableOrderNotFound"
	case CashDeskIsOpen:
		return "CashDeskIsOpen"
	case OriginTableIsFree:
		return "OriginTableIsFree"
	case DestinationTableIsReserved:
		return "DestinationTableIsReserved"
	case CantChangeSubProduct:
		return "CantChangeSubProduct"
	case UsernameExist:
		return "UsernameExist"
	case SumValuesNotEqual:
		return "SumValuesNotEqual"
	case CantIncreaseShouldNewOrder:
		return "CantIncreaseShouldNewOrder"
	case CostumerNotSelected:
		return "CostumerNotSelected"
	}
	return ""
}

func (s ServerText) FromString(e string) ServerText {
	switch e {
	case "NotFound":
		return NotFound
	case "TryAgain":
		return TryAgain
	case "InputTimeInvalid":
		return InputTimeInvalid
	case "InvalidInput":
		return InvalidInput
	case "IsExistWorkTime":
		return IsExistWorkTime
	case "UserNotFound":
		return UserNotFound
	case "AddressNotFound":
		return AddressNotFound
	case "WorkTimeNotFound":
		return WorkTimeNotFound
	case "WorkTimeOutOfNumber":
		return WorkTimeOutOfNumber
	case "WorkTimeFilled":
		return WorkTimeFilled
	case "NotEnoughChar":
		return NotEnoughChar
	case "NotReceivedFile":
		return NotReceivedFile
	case "NotAcceptable":
		return NotAcceptable
	case "DateIsNotFree":
		return DateIsNotFree
	case "UserOrPassInvalid":
		return UserOrPassInvalid
	case "Forbidden":
		return Forbidden
	case "InputBirthdayInvalid":
		return InputBirthdayInvalid
	case "InputNationalCodeInvalid":
		return InputNationalCodeInvalid
	case "BadRequest":
		return BadRequest
	case "ExceedLimit":
		return ExceedLimit
	case "TitleAddressInvalid":
		return TitleAddressInvalid
	case "InvalidLocation":
		return InvalidLocation
	case "InvalidTitleInput":
		return InvalidTitleInput
	case "ExistDiscountInTime":
		return ExistDiscountInTime
	case "TitleCharIsNotEnough":
		return TitleCharIsNotEnough
	case "PriorityInvalid":
		return PriorityInvalid
	case "WebsiteInvalid":
		return WebsiteInvalid
	case "WebsiteDuplicated":
		return WebsiteDuplicated
	case "InputDuplicated":
		return InputDuplicated
	case "SessionExist":
		return SessionExist
	case "SessionSelected":
		return SessionSelected
	case "SessionNotFound":
		return SessionNotFound
	case "SessionFilled":
		return SessionFilled
	case "SessionReserved":
		return SessionReserved
	case "OverflowDay":
		return OverflowDay
	case "SessionNotAvailable":
		return SessionNotAvailable
	case "SessionFree":
		return SessionFree
	case "SelectJustOneSession":
		return SelectJustOneSession
	case "SessionInPayment":
		return SessionInPayment
	case "SessionDisabled":
		return SessionDisabled
	case "InvalidPhoneNumber":
		return InvalidPhoneNumber
	case "TimePast":
		return TimePast
	case "TransactionNotFound":
		return TransactionNotFound
	case "TransactionExpired":
		return TransactionExpired
	case "TransactionFailed":
		return TransactionFailed
	case "OrderNotFound":
		return OrderNotFound
	case "ShopClosed":
		return ShopClosed
	case "DeliveryDisabled":
		return DeliveryDisabled
	case "OutOfServiceArea":
		return OutOfServiceArea
	case "CarInParking":
		return CarInParking
	case "ProductNotFound":
		return ProductNotFound
	case "ShopNotFound":
		return ShopNotFound
	case "TableNotFound":
		return TableNotFound
	case "TableReserved":
		return TableReserved
	case "ShouldOnePayment":
		return ShouldOnePayment
	case "TicketNotFound":
		return TicketNotFound
	case "PaymentPaid":
		return PaymentPaid
	case "OrderNotClosed":
		return OrderNotClosed
	case "MaxLowerThanMin":
		return MaxLowerThanMin
	case "TableHasNotPaidOrders":
		return TableHasNotPaidOrders
	case "CantChangePaidOrders":
		return CantChangePaidOrders
	case "RequestTypeNotFound":
		return RequestTypeNotFound
	case "TableOrderNotFound":
		return TableOrderNotFound
	case "CashDeskIsOpen":
		return CashDeskIsOpen
	case "OriginTableIsFree":
		return OriginTableIsFree
	case "DestinationTableIsReserved":
		return DestinationTableIsReserved
	case "CantChangeSubProduct":
		return CantChangeSubProduct
	case "UsernameExist":
		return UsernameExist
	case "SumValuesNotEqual":
		return SumValuesNotEqual
	case "CantIncreaseShouldNewOrder":
		return CantIncreaseShouldNewOrder
	case "CostumerNotSelected":
		return CostumerNotSelected
	}
	return NotDefined
}

const (
	NotFound ServerText = iota
	TryAgain
	InputTimeInvalid
	InvalidInput
	IsExistWorkTime
	UserNotFound
	AddressNotFound
	WorkTimeNotFound
	WorkTimeOutOfNumber
	WorkTimeFilled
	NotEnoughChar
	NotReceivedFile
	NotAcceptable
	DateIsNotFree
	UserOrPassInvalid
	Forbidden
	InputBirthdayInvalid
	InputNationalCodeInvalid
	BadRequest
	ExceedLimit
	TitleAddressInvalid
	InvalidLocation
	NotDefined
	InvalidTitleInput
	ExistDiscountInTime
	TitleCharIsNotEnough
	PriorityInvalid
	RefineryNumberInvalid
	WebsiteInvalid
	WebsiteDuplicated
	InputDuplicated
	SessionExist
	SessionSelected
	SessionNotFound
	SessionFilled
	SessionReserved
	OverflowDay
	SessionNotAvailable
	SessionFree
	SelectJustOneSession
	SessionInPayment
	SessionDisabled
	InvalidPhoneNumber
	TimePast
	TransactionNotFound
	TransactionExpired
	TransactionFailed
	OrderNotFound
	ShopClosed
	DeliveryDisabled
	OutOfServiceArea
	CarInParking
	ProductNotFound
	ShopNotFound
	TableNotFound
	TableReserved
	ShouldOnePayment
	TicketNotFound
	PaymentPaid
	OrderNotClosed
	MaxLowerThanMin
	TableHasNotPaidOrders
	CantChangePaidOrders
	RequestTypeNotFound
	TableOrderNotFound
	CashDeskIsOpen
	OriginTableIsFree
	DestinationTableIsReserved
	CantChangeSubProduct
	UsernameExist
	SumValuesNotEqual
	CantIncreaseShouldNewOrder
	CostumerNotSelected
)
