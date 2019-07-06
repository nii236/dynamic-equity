package contribution

// Multiplier is used to scale different types of contributions
type Multiplier int

// CashMultiplier to calculate value of contribution
const CashMultiplier = 4

// NonCashMultiplier to calculate value of contribution
const NonCashMultiplier = 2

// Type of contribution
type Type string

// CashType of contribution
const CashType Type = "cash"

// NonCashType of contribution
const NonCashType Type = "non-cash"

// Amount of contribution
func Amount(contribType Type, multiplier Multiplier, amount int) int {
	return 0
}
