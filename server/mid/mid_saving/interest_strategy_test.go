package main

import (
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"math"
	"reflect"
	"testing"
)

func TestParsing(t *testing.T) {
	acc := pb.SavingAccount{
		Id:          "0011-abcd-e123",
		UserID:      "1234-56ab-ed23",
		Balance:     1234567892,
		TermType:    "MONTHS",
		Term:        3,
		CreatedDate: "01012024",
		DueDate:     "01042024",
		Rate:        0.045,
	}
	accPT := &SavingAccountPT{}
	accPT.ParseFrom(&acc)
	if accPT.Id != acc.Id || accPT.UserID != acc.UserID || accPT.Balance != acc.Balance || accPT.TermType != acc.TermType || accPT.Term != acc.Term || accPT.CreatedDate != acc.CreatedDate || accPT.DueDate != acc.DueDate || accPT.InterestRate != acc.Rate {
		t.Errorf("All fields are expected the same, but result having difference")
	}
}

func TestGetCalculator(t *testing.T) {
	accPT := &SavingAccountPT{
		DueDate: "05072024",
	}
	withdrawDate := "11062024"

	accPT.GetCalculator(withdrawDate)
	if reflect.TypeOf(accPT.Calculator).Elem() != reflect.TypeOf(EarlyInterestCalculator{}) {
		t.Errorf("Reference expected %v but the result is %v", reflect.TypeOf(EarlyInterestCalculator{}), reflect.TypeOf(accPT.Calculator).Elem())
	}

	withdrawDate = "02082024"
	accPT.GetCalculator(withdrawDate)
	if reflect.TypeOf(accPT.Calculator).Elem() != reflect.TypeOf(OnTimeInterestCalculator{}) {
		t.Errorf("Reference expected %v but the result is %v", reflect.TypeOf(OnTimeInterestCalculator{}), reflect.TypeOf(accPT.Calculator).Elem())
	}
}

func TestTermInDays(t *testing.T) {
	acc1 := SavingAccountPT{TermType: "DAYS", Term: 21}
	output := acc1.TermInDays()
	expected := int32(21)
	if output != expected {
		t.Errorf("Expected %v but result is %v", expected, output)
	}

	acc2 := SavingAccountPT{TermType: "MONTHS", Term: 2}
	output = acc2.TermInDays()
	expected = int32(60)
	if output != expected {
		t.Errorf("Expected %v but result is %v", expected, output)
	}

	acc3 := SavingAccountPT{TermType: "YEARS", Term: 5}
	output = acc3.TermInDays()
	expected = int32(1800)
	if output != expected {
		t.Errorf("Expected %v but result is %v", expected, output)
	}

	// NOT REACHED
	acc4 := SavingAccountPT{TermType: "DECADES", Term: 2}
	output = acc4.TermInDays()
	expected = int32(2)
	if output != expected {
		t.Errorf("Expected %v but result is %v", expected, output)
	}
}

func TestCalculateRateOnTime(t *testing.T) {
	acc := &SavingAccountPT{
		InterestRate: 0.04,
		TermType:     "MONTHS",
		Term:         3,
		CreatedDate:  "01062024",
		DueDate:      "30082024",
		Calculator:   &OnTimeInterestCalculator{},
	}
	withdrawDate := "05092024"
	output := acc.Calculator.CalculateRate(acc, withdrawDate)
	expected := 0.0102
	if math.Abs(output/float64(expected)) > 1.05 || math.Abs(output/float64(expected)) < 0.95 {
		t.Errorf("Expected %v but output %v", expected, output)
	}
}

func TestCalculateRateEarly(t *testing.T) {
	acc := &SavingAccountPT{
		InterestRate: 0.04,
		TermType:     "MONTHS",
		Term:         3,
		CreatedDate:  "01062024",
		DueDate:      "30082024",
		Calculator:   &EarlyInterestCalculator{},
	}
	withdrawDate := "01072024"
	output := acc.Calculator.CalculateRate(acc, withdrawDate)
	expected := 0.000833
	if math.Abs(output/float64(expected)) > 1.05 || math.Abs(output/float64(expected)) < 0.95 {
		t.Errorf("Expected %v but output %v", expected, output)
	}
}

func TestDelegatingCalcInterest(t *testing.T) {
	acc := &SavingAccountPT{
		InterestRate: 0.04,
		TermType:     "MONTHS",
		Term:         3,
		CreatedDate:  "01062024",
		DueDate:      "30082024",
		Calculator:   &EarlyInterestCalculator{},
	}
	withdrawDate := "01082024"
	output1 := acc.CalculateRate(withdrawDate)
	output2 := acc.Calculator.CalculateRate(acc, withdrawDate)
	epsilon := 0.000001
	if math.Abs(output1-output2) > epsilon {
		t.Errorf("Result1 %v but result2 %v", output1, output2)
	}
}
