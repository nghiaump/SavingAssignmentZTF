package main

import (
	"testing"
)

func TestCalculateDueDateDays(t *testing.T) {
	createdDate := "01012022"
	expectedDueDate := "22012022"
	termType := "DAYS"
	term := 21

	result := CalculateDueDate(termType, term, createdDate)

	if result != expectedDueDate {
		t.Errorf("Expected due date %s, but output is %s", expectedDueDate, result)
	}
}

func TestCalculateDueDateMonths(t *testing.T) {
	createdDate := "01012022"
	expectedDueDate := "31012022"
	termType := "MONTHS"
	term := 1

	result := CalculateDueDate(termType, term, createdDate)

	if result != expectedDueDate {
		t.Errorf("Expected due date %s, but output is %s", expectedDueDate, result)
	}
}

func TestCalculateDueDateYears(t *testing.T) {
	createdDate := "01012022"
	expectedDueDate := "27122022"
	termType := "YEARS"
	term := 1

	result := CalculateDueDate(termType, term, createdDate)

	if result != expectedDueDate {
		t.Errorf("Expected due date %s, but output is %s", expectedDueDate, result)
	}
}

func TestCalculateDueDateErr(t *testing.T) {
	createdDate := "0a1012022"
	expectedDueDate := ""
	termType := "DAYS"
	term := 21

	result := CalculateDueDate(termType, term, createdDate)

	if result != expectedDueDate {
		t.Errorf("Expected due date %s, but output is %s", expectedDueDate, result)
	}
}

func TestFindFixedInterestRate(t *testing.T) {
	termType := "MONTHS"
	term := int32(6)
	kycLevel := int32(2)

	expectedRate := float32(0.05)
	result := FindFixedInterestRate(termType, term, kycLevel)

	if result != expectedRate {
		t.Errorf("Expected interest rate %f, but output is %f", expectedRate, result)
	}
}

func TestLaterThanTrue(t *testing.T) {
	withdrawnDate := "15012022"
	dueDate := "01012022"

	result := LaterThan(withdrawnDate, dueDate)

	if !result {
		t.Error("Expected true, but output is false")
	}
}

func TestLaterThanFalse(t *testing.T) {
	withdrawnDate := "01012022"
	dueDate := "15012022"

	result := LaterThan(withdrawnDate, dueDate)

	if result {
		t.Error("Expected false, but output is true")
	}
}

func TestLaterThanError(t *testing.T) {
	withdrawnDate := "0112022"
	dueDate := "15012022"

	result := LaterThan(withdrawnDate, dueDate)

	if result {
		t.Error("Expected false, but output is true")
	}
}

func TestCalculateOnTimeInterestDays(t *testing.T) {
	termType := "DAYS"
	term := int32(30)
	fixedRate := float32(0.03)

	expectedInterest := float32(0.0025)
	result := CalculateOnTimeInterest(termType, term, fixedRate)

	rate := float64(result) / float64(expectedInterest)
	if rate > 1.05 || rate < 0.95 {
		t.Errorf("Expected interest %f, but output is %f", expectedInterest, result)
	}
}

func TestCalculateOnTimeInterestMonths(t *testing.T) {
	termType := "MONTHS"
	term := int32(6)
	fixedRate := float32(0.05)

	expectedInterest := float32(0.025)
	result := CalculateOnTimeInterest(termType, term, fixedRate)

	if result != expectedInterest {
		t.Errorf("Expected interest %f, but output is %f", expectedInterest, result)
	}
}

func TestCalculateOnTimeInterestYears(t *testing.T) {
	termType := "YEARS"
	term := int32(2)
	fixedRate := float32(0.06)

	expectedInterest := float32(0.12)
	result := CalculateOnTimeInterest(termType, term, fixedRate)

	if result != expectedInterest {
		t.Errorf("Expected interest %f, but output is %f", expectedInterest, result)
	}
}

func TestCalculateOnTimeInterestDefault(t *testing.T) {
	termType := "YEARSSSSS"
	term := int32(2)
	fixedRate := float32(0.06)

	expectedInterest := float32(0.07)
	result := CalculateOnTimeInterest(termType, term, fixedRate)

	if result != expectedInterest {
		t.Errorf("Expected interest %f, but output is %f", expectedInterest, result)
	}
}

func TestCalculatePassedDays(t *testing.T) {
	createdDate := "01012022"
	withdrawnDate := "15012022"

	expectedDays := 14
	result := CalculatePassedDays(createdDate, withdrawnDate)

	if result != expectedDays {
		t.Errorf("Expected passed days %d, but output is %d", expectedDays, result)
	}
}
