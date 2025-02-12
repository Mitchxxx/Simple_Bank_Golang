package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
	GBP = "GBP"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD, GBP:
		return true
	}
	return false
}