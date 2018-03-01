package main

import (
	"testing"
	"time"
)

func CompareGroundTruth(t *testing.T, expected, actual PlanEntry) {
	if expected != actual {
		t.Error("Expected: ", expected, "got:", actual)
	}
}

func TestGetRepaymentPlan(t *testing.T) {
	startDate := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	duration := 2 * 12
	paymentPlan := GetRepaymentPlan(duration, 0.05, 5000, startDate)

	if len(paymentPlan) != duration {
		t.Error("Expected: ", duration, "got:", len(paymentPlan))
	}

	expectedFirstEntry := PlanEntry{startDate, 21936, 19853, 2083,
		500000, 480147}
	CompareGroundTruth(t, expectedFirstEntry, paymentPlan[0])

	expectedLastEntry := PlanEntry{startDate.AddDate(0, duration-1, 0), 21928,
		21837, 91, 21837, 0}
	CompareGroundTruth(t, expectedLastEntry, paymentPlan[len(paymentPlan)-1])
}

func TestCalculateAnnuity(t *testing.T) {
	duration := 2 * 12
	ratePerPeriod := 0.05 / 12.0
	loanAmount := 5000.0

	annuity := CalculateAnnuity(duration, ratePerPeriod, loanAmount)
	expected_annuity := 21936
	if annuity != expected_annuity {
		t.Error("Expected: ", expected_annuity, "got:", annuity)
	}
}
