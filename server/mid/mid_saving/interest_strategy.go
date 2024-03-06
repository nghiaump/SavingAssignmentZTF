package main

import (
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
)

// SavingAccountPT includes Calculator, which uses StrategyPattern
type SavingAccountPT struct {
	Id           string
	UserID       string
	Balance      int64
	TermInDays   int32
	CreatedDate  string
	DueDate      string
	InterestRate float32
	Calculator   IInterestCalculator
}

func (acc *SavingAccountPT) ParseFrom(pbAcc *pb.SavingAccount) {
	acc.Id = pbAcc.GetId()
	acc.UserID = pbAcc.GetUserId()
	acc.Balance = pbAcc.GetBalance()
	acc.TermInDays = pbAcc.TermInDays
	acc.CreatedDate, _ = ConvertFromISO8601(pbAcc.CreatedDate)
	acc.DueDate, _ = ConvertFromISO8601(pbAcc.DueDate)
	acc.InterestRate = pbAcc.GetRate()
}

func (acc *SavingAccountPT) GetCalculator(withDrawDate string) {
	if LaterThan(withDrawDate, acc.DueDate) {
		log.Println("OnTimeInterest")
		acc.Calculator = &OnTimeInterestCalculator{}
	} else {
		log.Println("EarlyInterest")
		acc.Calculator = &EarlyInterestCalculator{}
	}
}

func (acc *SavingAccountPT) CalculateRate(withdrawnDate string) float64 {
	return acc.Calculator.CalculateRate(acc, withdrawnDate)
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
	log.Printf("rate: %v\n", float64(acc.TermInDays)/float64(360)*float64(acc.InterestRate))
	return float64(acc.TermInDays) / float64(360) * float64(acc.InterestRate)
}
