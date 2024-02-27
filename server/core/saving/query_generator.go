package main

import (
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
)

const NilFlagKYCFilter = -1
const NilFlagTermFilter = -1
const NilFlagDueDateFilter = "-1"
const NilFlagMinBalanceFilter = -1

func GenerateQuery(filterObj *pb.Filter) map[string]interface{} {
	var allFilters []map[string]interface{}

	kycFilter := FilterByKYC(filterObj)
	termFilter := FilterByTerm(filterObj)
	dueDateRangeFilter := FilterByDueDateRange(filterObj)
	minBalanceFilter := FilterByMinBalance(filterObj)

	if kycFilter != nil {
		allFilters = append(allFilters, kycFilter)
	}

	if termFilter != nil {
		allFilters = append(allFilters, termFilter)
	}

	if dueDateRangeFilter != nil {
		allFilters = append(allFilters, dueDateRangeFilter)
	}

	if minBalanceFilter != nil {
		allFilters = append(allFilters, minBalanceFilter)
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": allFilters,
			},
		},
	}
	log.Printf("final query: %v\n", query)
	return query
}

func FilterByKYC(filterObj *pb.Filter) map[string]interface{} {
	if filterObj.Kyc == NilFlagKYCFilter {
		return nil
	}

	kycQuery := map[string]interface{}{
		"term": map[string]interface{}{
			"kyc": filterObj.Kyc,
		},
	}

	return kycQuery
}

func FilterByTerm(filterObj *pb.Filter) map[string]interface{} {
	if filterObj.TermInDays == NilFlagTermFilter {
		return nil
	}

	termQuery := map[string]interface{}{
		"term": map[string]interface{}{
			"term_in_days": filterObj.TermInDays,
		},
	}

	return termQuery
}

func FilterByDueDateRange(filterObj *pb.Filter) map[string]interface{} {
	if filterObj.DueDateEarliest == NilFlagDueDateFilter {
		return nil
	}

	dueDateEarliest, _ := ConvertToISO8601(filterObj.DueDateEarliest)
	dueDateLatest, _ := ConvertToISO8601(filterObj.DueDateLatest)

	dueDateRangeQuery := map[string]interface{}{
		"range": map[string]interface{}{
			"due_date": map[string]interface{}{
				"gte": dueDateEarliest,
				"lte": dueDateLatest,
			},
		},
	}

	return dueDateRangeQuery
}

func FilterByMinBalance(filterObj *pb.Filter) map[string]interface{} {
	if filterObj.MinBalance == NilFlagMinBalanceFilter {
		return nil
	}

	minBalanceQuery := map[string]interface{}{
		"range": map[string]interface{}{
			"balance": map[string]interface{}{
				"gte": filterObj.MinBalance,
			},
		},
	}

	return minBalanceQuery
}
