package main

import (
	"fmt"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
	"time"
)

const DateLayout = "02012006"
const NilFlagInt = -1
const NilFlagString = "-1"

func CalculateDueDate(termType string, term int, createdDate string) string {
	startDate, err := time.Parse(DateLayout, createdDate)
	if err != nil {
		return ""
	}
	var dueTime time.Time

	switch termType {
	case "DAYS":
		dueTime = startDate.Add(time.Duration(term) * 24 * time.Hour)
	case "MONTHS":
		dueTime = startDate.Add(time.Duration(term) * 30 * 24 * time.Hour)
	case "YEARS":
		dueTime = startDate.Add(time.Duration(term) * 360 * 24 * time.Hour)
	}
	return dueTime.Format("02012006")
}

type FixedInterestRate struct {
	kycLevel int32
	termType string
	term     int32
	rate     float32
}

var FixedInterestRateList = []FixedInterestRate{
	{2, "DAYS", 21, 0.03},
	{2, "MONTHS", 3, 0.04},
	{2, "MONTHS", 6, 0.05},
	{2, "MONTHS", 12, 0.06},
	{3, "DAYS", 21, 0.035},
	{3, "MONTHS", 3, 0.045},
	{3, "MONTHS", 6, 0.055},
	{3, "MONTHS", 12, 0.065},
}

func FindFixedInterestRate(termType string, term int32, kycLevel int32) float32 {
	for _, value := range FixedInterestRateList {
		if value.kycLevel == kycLevel && value.termType == termType && value.term == term {
			return value.rate
		}
	}
	return 0.3 // NOT REACH, DEFAULT
}

func LaterThan(withdrawnDateStr string, dueDateStr string) bool {
	date1, err1 := time.Parse(DateLayout, withdrawnDateStr)
	date2, err2 := time.Parse(DateLayout, dueDateStr)
	if err1 != nil || err2 != nil {
		fmt.Println("Error while parsing dates")
		return false
	}
	if !date1.Before(date2) {
		return true
	}
	return false
}

func CalculateOnTimeInterest(termType string, term int32, fixedRate float32) float32 {
	switch termType {
	case "DAYS":
		return float32(term) / float32(360) * fixedRate
	case "MONTHS":
		return float32(term) / float32(12) * fixedRate
	case "YEARS":
		return float32(term) * fixedRate
	}
	return 0.07 // Dummy value
}

func CalculatePassedDays(createdDateStr string, withdrawnDateStr string) int {
	date1, err1 := time.Parse(DateLayout, createdDateStr)
	date2, err2 := time.Parse(DateLayout, withdrawnDateStr)
	if err1 != nil || err2 != nil {
		log.Printf("cannot parse Date")
		//return 0
	}
	return int(date2.Sub(date1).Hours() / 24)
}

func ConvertToISO8601(date string) (string, error) {
	parsedDate, err := time.Parse(DateLayout, date)
	if err != nil {
		return "", err
	}

	isoDate := parsedDate.Format("2006-01-02T15:04:05Z")
	return isoDate, nil
}

func ConvertFromISO8601(isoDate string) (string, error) {
	parsedDate, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return "", err
	}

	ddmmyyyyDate := parsedDate.Format(DateLayout)
	return ddmmyyyyDate, nil
}

func ValidateFilterRequest(req *pb.Filter) bool {
	if req.Kyc != NilFlagInt && req.Kyc != 2 && req.Kyc != 3 {
		return false
	}
	if req.TermInDays != NilFlagInt && req.TermInDays < 0 {
		return false
	}
	if req.DueDateEarliest != NilFlagString {
		if _, err := time.Parse(DateLayout, req.DueDateEarliest); err != nil {
			return false
		}
	}

	if req.DueDateLatest != NilFlagString {
		if _, err := time.Parse(DateLayout, req.DueDateLatest); err != nil {
			return false
		}
	}

	if req.PageSize < 1 {
		return false
	}
	if req.PageIndex < 1 {
		return false
	}

	// Passed all constrains
	return true
}
