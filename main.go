package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func generatePlanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST is supported", http.StatusForbidden)
		return
	}

	// parse POST request payload
	decoder := json.NewDecoder(r.Body)
	var pr PlanRequest
	err := decoder.Decode(&pr)
	if err != nil {
		msg := fmt.Sprintf("Error: %s", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// process payload and check for errors
	nominalRate, err := strconv.ParseFloat(pr.NominalRate, 64)
	if err != nil {
		msg := fmt.Sprintf("Error: %s", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	nominalRate /= 100.0
	loanAmount, err := strconv.ParseFloat(pr.LoanAmount, 64)
	if err != nil {
		msg := fmt.Sprintf("Error: %s", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	if pr.Duration == 0 {
		msg := "Duration not set"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	if pr.StartDate.IsZero() {
		msg := "StartDate not set"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	paymentPlan := GetRepaymentPlan(pr.Duration, nominalRate, loanAmount, pr.StartDate)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentPlan)
}

func main() {
	http.HandleFunc("/generate-plan", generatePlanHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
