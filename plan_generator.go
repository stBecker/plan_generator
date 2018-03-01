package main

import (
	"encoding/json"
	"math"
	"strconv"
	"time"
)

const DaysPerMonth float64 = 30
const DaysPerYear float64 = 360

type PlanRequest struct {
	LoanAmount  string
	NominalRate string
	Duration    int
	StartDate   time.Time
}

type PlanEntry struct {
	// values are in cents
	Date                          time.Time
	BorrowerPaymentAmount         int
	Principal                     int
	Interest                      int
	InitialOutstandingPrincipal   int
	RemainingOutstandingPrincipal int
}

func (u *PlanEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Date                          time.Time
		BorrowerPaymentAmount         string
		Principal                     string
		Interest                      string
		InitialOutstandingPrincipal   string
		RemainingOutstandingPrincipal string
	}{
		Date: u.Date,
		BorrowerPaymentAmount:         CentsToEuro(u.BorrowerPaymentAmount),
		Principal:                     CentsToEuro(u.Principal),
		Interest:                      CentsToEuro(u.Interest),
		InitialOutstandingPrincipal:   CentsToEuro(u.InitialOutstandingPrincipal),
		RemainingOutstandingPrincipal: CentsToEuro(u.RemainingOutstandingPrincipal),
	})
}

func CentsToEuro(x int) string {
	return strconv.Itoa(x/100) + "." + strconv.Itoa(x%100)
}

func EuroToCents(x float64) int {
	return int(math.Round(x * 100.0))
}

func CalculateAnnuity(duration int, ratePerPeriod, loanAmount float64) int {
	denominator := 1 - math.Pow(1+ratePerPeriod, -float64(duration))
	annuityEuro := ratePerPeriod * loanAmount / denominator
	annuity := EuroToCents(annuityEuro)
	return annuity
}

func GetRepaymentPlan(duration int, nominalRate float64, loanAmount float64, startDate time.Time) []PlanEntry {
	var paymentPlan []PlanEntry
	ratePerPeriod := nominalRate * DaysPerMonth / DaysPerYear
	annuity := CalculateAnnuity(duration, ratePerPeriod, loanAmount)

	for i := 0; i < duration; i++ {
		var entry PlanEntry

		if i == 0 {
			// set up first installment
			entry.InitialOutstandingPrincipal = EuroToCents(loanAmount)
			entry.Date = startDate
		} else {
			prevEntry := paymentPlan[i-1]
			entry.InitialOutstandingPrincipal = prevEntry.RemainingOutstandingPrincipal
			entry.Date = prevEntry.Date.AddDate(0, 1, 0)
		}

		entry.Interest = EuroToCents(ratePerPeriod * float64(entry.InitialOutstandingPrincipal) / 100.0)
		entry.Principal = annuity - entry.Interest

		// the current principal mustn't be higher than the outstanding principal
		if entry.Principal > entry.InitialOutstandingPrincipal {
			entry.Principal = entry.InitialOutstandingPrincipal
		} else {
			entry.Principal = annuity - entry.Interest
		}
		entry.RemainingOutstandingPrincipal = entry.InitialOutstandingPrincipal - entry.Principal
		entry.BorrowerPaymentAmount = entry.Principal + entry.Interest

		paymentPlan = append(paymentPlan, entry)
	}

	return paymentPlan
}
