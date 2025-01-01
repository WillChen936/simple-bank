package utils

const (
	USD string = "USD"
	EUR string = "EUR"
	CAD string = "CAD"
)

var Currencies []string

func init() {
	Currencies = []string{
		USD, EUR, CAD,
	}
}
