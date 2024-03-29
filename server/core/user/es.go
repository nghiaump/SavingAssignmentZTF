package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/golang/glog"
	"github.com/google/uuid"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"os"
	"reflect"
	"strings"
	"time"
)

func ConvertToISO8601(date string) (string, error) {
	parsedDate, err := time.Parse("02012006", date)
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

	ddmmyyyyDate := parsedDate.Format("02012006")
	return ddmmyyyyDate, nil
}

func CreateESClient() (*elasticsearch.Client, bool) {
	addressESContainer := os.Getenv(ContainerElasticSearchEnv)
	if addressESContainer == "" {
		glog.Info("ContainerElasticSearchEnv not found.")
		return nil, true
	}

	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://" + addressESContainer + ElasticSearchPort},
	})

	if err != nil {
		glog.Errorf("Failed to create Elasticsearch client: %v", err)
		return nil, true
	} else {
		glog.Info("Connected to ElasticSearch")
	}
	return esClient, false
}

func InitIndex(indexName string, esClient *elasticsearch.Client) {
	exist, _ := esClient.Indices.Exists([]string{indexName})
	if !(exist != nil && exist.StatusCode == 200) {
		glog.Info("InitIndex")
		_, err3 := esClient.Indices.Create(indexName)
		if err3 != nil {
			fmt.Println(err3)
		}
	} else {
		glog.Info(`InitIndex: Index "user" existing`)
	}
}

func CreateIndexingRequest(req interface{}, indexName string) esapi.IndexRequest {
	// in user, req is *pb.User
	out, _ := uuid.NewUUID()
	req.(*pb.User).Id = out.String()
	val := reflect.ValueOf(req)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		fmt.Println("Error: val.Kind() != reflect.Struct")
	}

	doc := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldName := field.Tag.Get(ESDocumentTag)

		if fieldName != "" {
			doc[fieldName] = val.Field(i).Interface()
		}
	}

	// Chuyển đổi map thành chuỗi JSON
	jsonStr, err := json.Marshal(doc)
	if err != nil {
		// TODO
	}

	glog.Info("Test json marshal %v", string(jsonStr))
	indexReq := esapi.IndexRequest{
		Index:   indexName,
		Body:    strings.NewReader(string(jsonStr)),
		Refresh: "true",
	}

	return indexReq
}

func SearchOneUserByUniqueTextField(fieldName string, value string, client *elasticsearch.Client) *pb.User {
	var r map[string]interface{}
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": FilterByStringExact(fieldName, value),
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		glog.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(ESUserIndex),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		glog.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			glog.Fatalf("SearchOneUserByUniqueTextField: Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			glog.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		glog.Fatalf("SearchOneUserByUniqueTextField: Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	// glog.Infof(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		// Convert the bytes data to JSON
		jsonData, err := json.Marshal(doc)
		if err != nil {
			glog.Errorf("SearchOneUserByUniqueTextField: Error marshaling JSON: %v", err)
		}

		// Convert JSON to struct
		userObj := pb.User{}
		if err := json.Unmarshal(jsonData, &userObj); err != nil {
			glog.Infof("SearchOneUserByUniqueTextField: Error unmarshalling document in response: %v", err)
		} else {
			glog.Infof("SearchOneUserByUniqueTextField: Unmarshaled successfully: %v", userObj)
			return &userObj
		}
	}

	return nil
}

func SearchUsersByFiltersHelper(filterObj *pb.UserFilter, client *elasticsearch.Client) []*pb.User {
	query := GenerateQuery(filterObj)
	var r map[string]interface{}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		glog.Fatalf("SearchUsersByFiltersHelper: Fail to encode query: %s", err)
	}

	// Perform the search request.
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(ESUserIndex),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		glog.Fatalf("SearchUsersByFiltersHelper: Failed to get response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			glog.Fatalf("SearchUsersByFiltersHelper: Failed to parse response body: %s", err)
		} else {
			// Print the response status and error information.
			glog.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		glog.Fatalf("SearchUsersByFiltersHelper: Failed to parse response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	//glog.Infof(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)

	userList := []*pb.User{}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		//glog.Info(" * ID=%s:\n", hit.(map[string]interface{})["_id"])
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		// Convert the bytes data to JSON
		jsonData, err := json.Marshal(doc)
		if err != nil {
			fmt.Println("SearchUsersByFiltersHelper: Error marshaling JSON:", err)
		}

		// Convert JSON to struct
		userObj := pb.User{}
		if err := json.Unmarshal(jsonData, &userObj); err != nil {
			glog.Infof("SearchUsersByFiltersHelper: Error unmarshalling document in response: %v", err)
		} else {
			glog.Infof("SearchUsersByFiltersHelper: after unmarshaled: %v", userObj)
			userList = append(userList, &userObj)
		}
	}

	return userList
}

const NilFlagInt = -1
const NilFlagString = "-1"
const MaxBalance = 999999999999

func GenerateQuery(filterObj *pb.UserFilter) map[string]interface{} {
	var allFilters []map[string]interface{}

	kycFilter := FilterByInt32Exact("kyc", filterObj.Kyc)
	registeredDateFilter := FilterByDateRange("registered_date", filterObj.RegisteredDateRangeEarliest, filterObj.RegisteredDateRangeLatest)
	genderFilter := FilterByInt32Exact("gender", filterObj.Gender)
	addressFilter := FilterByStringContained("address", filterObj.Address)

	if kycFilter != nil {
		allFilters = append(allFilters, kycFilter)
	}

	if registeredDateFilter != nil {
		allFilters = append(allFilters, registeredDateFilter)
	}

	if genderFilter != nil {
		allFilters = append(allFilters, genderFilter)
	}

	if addressFilter != nil {
		allFilters = append(allFilters, addressFilter)
	}

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": allFilters,
			},
		},
	}
	glog.Info("GenerateQuery: query: %v\n", query)
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
			//fieldName + ".keyword": searchText,
			fieldName: searchText,
		},
	}

	glog.Infof("FilterByStringExact: query %v\n", query)

	return query
}
