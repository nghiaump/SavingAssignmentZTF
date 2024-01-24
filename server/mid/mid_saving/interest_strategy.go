package main

import (
	"fmt"
	pb "github.com/nghiaump/savingaccproto"
)

// SavingAccountPT includes Calculator, which uses StrategyPattern
type SavingAccountPT struct {
	Id           string
	UserID       string
	Balance      int64
	TermType     string
	Term         int32
	CreatedDate  string
	DueDate      string
	InterestRate float32
	Calculator   IInterestCalculator
}

func (acc *SavingAccountPT) ParseFrom(pbAcc *pb.SavingAccount) {
	acc.Id = pbAcc.GetId()
	acc.UserID = pbAcc.GetUserID()
	acc.Balance = pbAcc.GetBalance()
	acc.TermType = pbAcc.GetTermType()
	acc.Term = pbAcc.GetTerm()
	acc.CreatedDate = pbAcc.GetCreatedDate()
	acc.DueDate = pbAcc.GetDueDate()
	acc.InterestRate = pbAcc.GetRate()
}

func (acc *SavingAccountPT) GetCalculator(withDrawDate string) {
	if LaterThan(withDrawDate, acc.DueDate) {
		acc.Calculator = &OnTimeInterestCalculator{}
	} else {
		acc.Calculator = &EarlyInterestCalculator{}
	}
}

func (acc *SavingAccountPT) CalculateRate(withdrawnDate string) float64 {
	return acc.Calculator.CalculateRate(acc, withdrawnDate)
}

func (acc *SavingAccountPT) TermInDays() int32 {
	switch acc.TermType {
	case "DAYS":
		return acc.Term
	case "MONTHS":
		return acc.Term * 30
	case "YEARS":
		return acc.Term * 360
	}
	return acc.Term // not reached
}

// IInterestCalculator is used for Strategy Pattern
type IInterestCalculator interface {
	CalculateRate(acc *SavingAccountPT, withdrawnDate string) float64
}

type EarlyInterestCalculator struct{}
type OnTimeInterestCalculator struct{}

const NonTermRate = 0.01

func (earlyInterest *EarlyInterestCalculator) CalculateRate(acc *SavingAccountPT, withdrawnDate string) float64 {
	passedDays := CalculatePassedDays(acc.CreatedDate, withdrawnDate)
	return float64(passedDays) / float64(360) * NonTermRate
}

func (onTimeInterest *OnTimeInterestCalculator) CalculateRate(acc *SavingAccountPT, withdrawnDate string) float64 {
	daysTerm := acc.TermInDays()
	fmt.Println(daysTerm)
	return float64(daysTerm) / float64(360) * float64(acc.InterestRate)
}
