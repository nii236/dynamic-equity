package contributions

import (
	"dynamic-equity/pie"
	"errors"
)

// Time contributed to the business
func Time(p *pie.Pie, hourlyRate int, hours int) (int, error) {
	return hourlyRate * hours * p.NonCashMultiplier, nil
}

// Expenses contributed to the business
func Expenses(p *pie.Pie, paid, reimbursed int) (int, error) {
	return (paid - reimbursed) * p.CashMultiplier, nil
}

// Sales contributed to the business
func Sales(p *pie.Pie, amount, royalty, cashPayment int) (int, error) {
	return ((amount * royalty) - cashPayment) * p.NonCashMultiplier, nil
}

// FindersFee contributed to the business
func FindersFee(p *pie.Pie, raised, cutoff int) (int, error) {
	return 0, errors.New("Not implemented")
}

// SuppliesNew contributed to the business (new purchase)
func SuppliesNew(p *pie.Pie, paid, reimbursed int) (int, error) {
	return (paid - reimbursed) * p.CashMultiplier, nil
}

// SuppliesRecent contributed to the business (less than a year old)
func SuppliesRecent(p *pie.Pie, paid, reimbursed int) (int, error) {
	return (paid - reimbursed) * p.NonCashMultiplier, nil
}

// SuppliesOld contributed to the business (over a year old)
func SuppliesOld(p *pie.Pie, fairMarketValue, reimbursed int) (int, error) {
	return (fairMarketValue - reimbursed) * p.NonCashMultiplier, nil
}

// EquipmentNew contributed to the business (new purchase)
func EquipmentNew(p *pie.Pie, paid, reimbursed int) (int, error) {
	return (paid - reimbursed) * p.CashMultiplier, nil
}

// EquipmentRecent contributed to the business (less than a year old)
func EquipmentRecent(p *pie.Pie, paid, reimbursed int) (int, error) {
	return (paid - reimbursed) * p.NonCashMultiplier, nil
}

// EquipmentOld contributed to the business (over a year old)
func EquipmentOld(p *pie.Pie, fairMarketValue, reimbursed int) (int, error) {
	return (fairMarketValue - reimbursed) * p.NonCashMultiplier, nil
}

// Facilities contributed to the business
func Facilities(p *pie.Pie, fairMarketValue, cashPayment int) (int, error) {
	return (fairMarketValue - cashPayment) * p.NonCashMultiplier, nil
}

// VehicleFuel contributed to the business
func VehicleFuel(p *pie.Pie, fuelCost, reimbursed int) (int, error) {
	return (fuelCost - reimbursed) * p.CashMultiplier, nil
}

// OtherCash contributed to the business
func OtherCash(p *pie.Pie, amount int) (int, error) {
	return amount * p.CashMultiplier, nil
}

// OtherNonCash contributed to the business
func OtherNonCash(p *pie.Pie, amount int) (int, error) {
	return amount * p.NonCashMultiplier, nil
}
