package main

import (
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
)

const NilFlagInt = -1
const NilFlagString = "-1"
const MaxBalance = 999999999999

func GenerateQuery(filterObj *pb.Filter) map[string]interface{} {
	var allFilters []map[string]interface{}

	kycFilter := FilterByInt32Exact(filterObj.Kyc, "kyc")
	termFilter := FilterByInt32Exact(filterObj.TermInDays, "term_in_days")
	dueDateRangeFilter := FilterByDateRange(filterObj.DueDateEarliest, filterObj.DueDateLatest, "due_date")
	minBalanceFilter := FilterByInt64Range(filterObj.MinBalance, MaxBalance, "balance")

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

func FilterByInt32Exact(value int32, fieldName string) map[string]interface{} {
	if value == NilFlagInt {
		return nil
	}

	query := map[string]interface{}{
		"term": map[string]interface{}{
			fieldName: value,
		},
	}

	return query
}

func FilterByInt32Range(valueMin int32, valueMax int32, fieldName string) map[string]interface{} {
	if valueMin == NilFlagInt {
		return nil
	}

	query := map[string]interface{}{
		"range": map[string]interface{}{
			fieldName: map[string]interface{}{
				"gte": valueMin,
				"lte": valueMax,
			},
		},
	}

	return query
}

func FilterByInt64Exact(value int64, fieldName string) map[string]interface{} {
	if value == NilFlagInt {
		return nil
	}

	query := map[string]interface{}{
		"term": map[string]interface{}{
			fieldName: value,
		},
	}

	return query
}

func FilterByInt64Range(valueMin int64, valueMax int64, fieldName string) map[string]interface{} {
	if valueMin == NilFlagInt {
		return nil
	}

	query := map[string]interface{}{
		"range": map[string]interface{}{
			fieldName: map[string]interface{}{
				"gte": valueMin,
				"lte": valueMax,
			},
		},
	}

	return query
}

func FilterByDateExact(dateString string, fieldName string) map[string]interface{} {
	if dateString == NilFlagString {
		return nil
	}

	date, _ := ConvertToISO8601(dateString)

	query := map[string]interface{}{
		"term": map[string]interface{}{
			fieldName: date,
		},
	}

	return query
}

func FilterByDateRange(dateEarliestString string, dateLatestString string, fieldName string) map[string]interface{} {
	if dateEarliestString == NilFlagString || dateLatestString == NilFlagString {
		return nil
	}

	dateEarliest, _ := ConvertToISO8601(dateEarliestString)
	dateLatest, _ := ConvertToISO8601(dateLatestString)

	query := map[string]interface{}{
		"range": map[string]interface{}{
			fieldName: map[string]interface{}{
				"gte": dateEarliest,
				"lte": dateLatest,
			},
		},
	}

	return query
}
