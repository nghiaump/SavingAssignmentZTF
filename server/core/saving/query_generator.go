package main

import (
	"github.com/golang/glog"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
)

const NilFlagInt = -1
const NilFlagString = "-1"
const MaxBalance = 999999999999

func GenerateQuery(filterObj *pb.Filter) map[string]interface{} {
	var allFilters []map[string]interface{}

	kycFilter := FilterByInt32Exact("kyc", filterObj.Kyc)
	termFilter := FilterByInt32Exact("term_in_days", filterObj.TermInDays)
	dueDateRangeFilter := FilterByDateRange("due_date", filterObj.DueDateEarliest, filterObj.DueDateLatest)
	minBalanceFilter := FilterByInt64Range("balance", filterObj.MinBalance, MaxBalance)

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
	glog.Infof("GenerateQuery: final query: %v\n", query)
	return query
}

func GenerateQueryWithAgg(filterObj *pb.Filter) map[string]interface{} {
	var allFilters []map[string]interface{}

	kycFilter := FilterByInt32Exact("kyc", filterObj.Kyc)
	termFilter := FilterByInt32Exact("term_in_days", filterObj.TermInDays)
	dueDateRangeFilter := FilterByDateRange("due_date", filterObj.DueDateEarliest, filterObj.DueDateLatest)
	minBalanceFilter := FilterByInt64Range("balance", filterObj.MinBalance, MaxBalance)

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
		"from": (filterObj.PageIndex - 1) * filterObj.PageSize,
		"size": filterObj.PageSize,

		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": allFilters,
			},
		},

		"aggs": map[string]interface{}{
			"total_balance": map[string]interface{}{
				"sum": map[string]interface{}{
					"field": "balance",
				},
			},
		},
	}

	glog.Infof("GenerateQueryWithAgg: final query: %v\n", query)
	return query
}

func FilterByInt32Exact(fieldName string, value int32) map[string]interface{} {
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

func FilterByInt32Range(fieldName string, valueMin int32, valueMax int32) map[string]interface{} {
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

func FilterByInt64Exact(fieldName string, value int64) map[string]interface{} {
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

func FilterByInt64Range(fieldName string, valueMin int64, valueMax int64) map[string]interface{} {
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

func FilterByDateExact(fieldName string, dateString string) map[string]interface{} {
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

func FilterByDateRange(fieldName string, dateEarliestString string, dateLatestString string) map[string]interface{} {
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

func FilterByStringContained(fieldName string, searchText string) map[string]interface{} {
	if searchText == NilFlagString {
		return nil
	}

	query := map[string]interface{}{
		"match": map[string]interface{}{
			fieldName: searchText,
		},
	}

	return query
}

func FilterByStringExact(fieldName string, searchText string) map[string]interface{} {
	if searchText == NilFlagString {
		return nil
	}

	query := map[string]interface{}{
		"term": map[string]interface{}{
			fieldName + ".keyword": searchText,
		},
	}

	return query
}
