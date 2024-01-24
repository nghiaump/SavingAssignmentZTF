package main

import (
	"testing"
)

func TestCreateSavingAccountPT_KYC2_Days(t *testing.T) {
	factory := &SavingFactoryKYC2{}
	account := factory.CreateSavingAccountPT("DAYS", 21)

	expectedRate := float32(0.03)
	if account.InterestRate != expectedRate {
		t.Errorf("Expected interest rate %f, but got %f", expectedRate, account.InterestRate)
	}
}

func TestCreateSavingAccountPT_KYC2_Months3(t *testing.T) {
	factory := &SavingFactoryKYC2{}
	account := factory.CreateSavingAccountPT("MONTHS", 3)

	expectedRate := float32(0.04)
	if account.InterestRate != expectedRate {
		t.Errorf("Expected interest rate %f, but got %f", expectedRate, account.InterestRate)
	}
}

func TestCreateSavingAccountPT_KYC2_Months6(t *testing.T) {
	factory := &SavingFactoryKYC2{}
	account := factory.CreateSavingAccountPT("MONTHS", 6)

	expectedRate := float32(0.05)
	if account.InterestRate != expectedRate {
		t.Errorf("Expected interest rate %f, but got %f", expectedRate, account.InterestRate)
	}
}

func TestCreateSavingAccountPT_KYC2_Months12(t *testing.T) {
	factory := &SavingFactoryKYC2{}
	account := factory.CreateSavingAccountPT("MONTHS", 12)

	expectedRate := float32(0.06)
	if account.InterestRate != expectedRate {
		t.Errorf("Expected interest rate %f, but got %f", expectedRate, account.InterestRate)
	}
}

func TestCreateSavingAccountPT_KYC3_Days(t *testing.T) {
	factory := &SavingFactoryKYC3{}
	account := factory.CreateSavingAccountPT("DAYS", 21)

	expectedRate := float32(0.035)
	if account.InterestRate != expectedRate {
		t.Errorf("Expected interest rate %f, but got %f", expectedRate, account.InterestRate)
	}
}

func TestCreateSavingAccountPT_KYC3_Months3(t *testing.T) {
	factory := &SavingFactoryKYC3{}
	account := factory.CreateSavingAccountPT("MONTHS", 3)

	expectedRate := float32(0.045)
	if account.InterestRate != expectedRate {
		t.Errorf("Expected interest rate %f, but got %f", expectedRate, account.InterestRate)
	}
}

func TestCreateSavingAccountPT_KYC3_Months6(t *testing.T) {
	factory := &SavingFactoryKYC3{}
	account := factory.CreateSavingAccountPT("MONTHS", 6)

	expectedRate := float32(0.055)
	if account.InterestRate != expectedRate {
		t.Errorf("Expected interest rate %f, but got %f", expectedRate, account.InterestRate)
	}
}

func TestCreateSavingAccountPT_KYC3_Months12(t *testing.T) {
	factory := &SavingFactoryKYC3{}
	account := factory.CreateSavingAccountPT("MONTHS", 12)

	expectedRate := float32(0.065)
	if account.InterestRate != expectedRate {
		t.Errorf("Expected interest rate %f, but got %f", expectedRate, account.InterestRate)
	}
}
