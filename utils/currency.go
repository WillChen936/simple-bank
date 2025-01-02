package utils

var Currencies = map[string]struct{}{
	"USD": {},
	"EUR": {},
	"CAD": {},
}

func IsSupportedCurrency(currency string) bool {
	_, exists := Currencies[currency]
	return exists
}
