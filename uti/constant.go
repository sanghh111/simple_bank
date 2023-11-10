package uti

const (
	DateTimeLayout = "2006-01-02 15:04:05"
)

// ERROR
const (
	Success                   = "00"
	RequestInfoEmpty          = "01"
	LangCodeEmpty             = "02"
	RequestTime               = "03"
	TransferSameAccount       = "04"
	FromAccountExisted        = "05"
	ToAccountExisted          = "06"
	FromAccountNotEnoughMoney = "07"
)

var MessInputError = map[string]string{
	RequestInfoEmpty:          "Request info Empty",
	LangCodeEmpty:             "Lang code Empty",
	RequestTime:               "Request Time Empty",
	Success:                   "Success",
	TransferSameAccount:       "Cannot transfer money to the same account",
	FromAccountExisted:        "From account not existed",
	ToAccountExisted:          "To account not existed",
	FromAccountNotEnoughMoney: "From account not enough money",
}
