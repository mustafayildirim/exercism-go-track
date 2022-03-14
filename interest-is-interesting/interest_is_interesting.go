package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance >= 5000 {
		return 2.475
	} else if balance >= 1000 {
		return 1.621
	} else if balance >= 0 {
		return 0.5
	}

	return 3.213
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	interestRate := InterestRate(balance)
	return balance * float64(interestRate) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	interest := Interest(balance)
	return balance + interest
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	annualBalance := AnnualBalanceUpdate(balance)
	yearsPassed := 1

	for annualBalance < targetBalance {
		annualBalance = AnnualBalanceUpdate(annualBalance)
		yearsPassed++
	}

	return yearsPassed
}
