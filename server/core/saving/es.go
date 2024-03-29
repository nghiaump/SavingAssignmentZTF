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
		glog.Info("ContainerElasticSearchEnv not found")
		return nil, true
	}

	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://" + addressESContainer + ElasticSearchPort},
	})

	if err != nil {
		glog.Infof("Error creating Elasticsearch client:", err)
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
		glog.Info(`InitIndes: "saving" existing`)
	}
}

func CreateIndexingRequest(req interface{}, indexName string) esapi.IndexRequest {
	//req *pb.SavingAccount
	out, _ := uuid.NewUUID()
	req.(*pb.SavingAccount).Id = out.String()
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

	glog.Infof("Test json marshal %v", string(jsonStr))
	indexReq := esapi.IndexRequest{
		Index:   indexName,
		Body:    strings.NewReader(string(jsonStr)),
		Refresh: "true",
	}

	return indexReq
}

func CreateUpdateRequest(req interface{}, indexName string) esapi.UpdateRequest {
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

	jsonStr, err := json.Marshal(doc)
	if err != nil {
		// TODO
	}

	glog.Infof("Test json marshal %v", string(jsonStr))
	updateReq := esapi.UpdateRequest{
		Index:   indexName,
		Body:    strings.NewReader(string(jsonStr)),
		Refresh: "true",
	}

	return updateReq
}

func SearchOneAccountByUniqueTextField(fieldName string, value string, client *elasticsearch.Client) *pb.SavingAccount {
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
		client.Search.WithIndex(ESSavingIndex),
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
			glog.Fatalf("Error parsing the response body: %s", err)
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
		glog.Fatalf("Error parsing the response body: %s", err)
	}
	//// Print the response status, number of results, and request duration.
	//glog.Infof(
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
			fmt.Println("Error marshaling JSON:", err)
		}

		// Convert JSON to struct
		accObj := pb.SavingAccount{}
		if err := json.Unmarshal(jsonData, &accObj); err != nil {
			glog.Infof("Error unmarshalling document in response:", err)
		} else {
			glog.Infof("Unmarshaled successfully: %v", accObj)
			return &accObj
		}
	}

	return nil
}

func SearchDocIDByUniqueTextField(fieldName string, value string, client *elasticsearch.Client) string {
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
		client.Search.WithIndex(ESSavingIndex),
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
			glog.Fatalf("Error parsing the response body: %s", err)
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
		glog.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	//glog.Infof(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		docID := hit.(map[string]interface{})["_id"].(string)
		return docID
	}

	return ""
}

func GetAllAccountsByUserIDHelper(userID string, client *elasticsearch.Client) []*pb.SavingAccount {
	var r map[string]interface{}
	var buf bytes.Buffer
	// match -> gan giong
	// term -> chinh xac giong
	query := map[string]interface{}{
		"query": FilterByStringExact("user_id", userID),
	}

	glog.Infof("GetAllAccountsByUserIDHelper: query: %v", query)

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		glog.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(ESSavingIndex),
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
			glog.Fatalf("Error parsing the response body: %s", err)
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
		glog.Fatalf("GetAllAccountsByUserIDHelper: Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	//glog.Infof(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)

	accList := []*pb.SavingAccount{}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		//glog.Info(" * ID=%s:\n", hit.(map[string]interface{})["_id"])
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		// Convert the bytes data to JSON
		jsonData, err := json.Marshal(doc)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
		}

		// Convert JSON to struct
		accStruct := pb.SavingAccount{}
		if err := json.Unmarshal(jsonData, &accStruct); err != nil {
			glog.Infof("GetAllAccountsByUserIDHelper: Error unmarshalling document in response:", err)
		} else {
			accList = append(accList, &accStruct)
		}
	}

	//glog.Infof(strings.Repeat("=", 37))
	return accList
}

func SearchAccountsByFiltersHelper(filterObj *pb.Filter, client *elasticsearch.Client) []*pb.SavingAccount {
	query := GenerateQuery(filterObj)
	var r map[string]interface{}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		glog.Errorf("SearchAccountsByFiltersHelper: Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(ESSavingIndex),
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
			glog.Fatalf("Error parsing the response body: %s", err)
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
		glog.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	glog.Info(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	accList := []*pb.SavingAccount{}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		//glog.Info(" * ID=%s:\n", hit.(map[string]interface{})["_id"])
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		// Convert the bytes data to JSON
		jsonData, err := json.Marshal(doc)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
		}

		// Convert JSON to struct
		accStruct := pb.SavingAccount{}
		if err := json.Unmarshal(jsonData, &accStruct); err != nil {
			glog.Info("Error unmarshalling document in response:", err)
		} else {
			accList = append(accList, &accStruct)
		}
	}

	//glog.Infof(strings.Repeat("=", 37))
	return accList
}

func SearchAccountsByFiltersWithPaging(filterObj *pb.Filter, client *elasticsearch.Client) ([]*pb.SavingAccount, int64, int64) {
	query := GenerateQueryWithAgg(filterObj)
	var r map[string]interface{}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		glog.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex(ESSavingIndex),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		glog.Fatalf("SearchAccountsByFiltersWithPaging: Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			glog.Fatalf("SearchAccountsByFiltersWithPaging: Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			glog.Errorf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		glog.Fatalf("SearchAccountsByFiltersWithPaging: Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	//glog.Infof(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)

	accList := []*pb.SavingAccount{}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		//glog.Info(" * ID=%s:\n", hit.(map[string]interface{})["_id"])
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		// Convert the bytes data to JSON
		jsonData, err := json.Marshal(doc)
		if err != nil {
			glog.Infof("SearchAccountsByFiltersWithPaging: Error marshaling JSON:", err)
		}

		// Convert JSON to struct
		accStruct := pb.SavingAccount{}
		if err := json.Unmarshal(jsonData, &accStruct); err != nil {
			glog.Infof("SearchAccountsByFiltersWithPaging: Error unmarshalling document in response:", err)
		} else {
			accList = append(accList, &accStruct)
		}
	}

	totalHits := int64(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	aggregations := r["aggregations"].(map[string]interface{})
	totalBalance := int64(aggregations["total_balance"].(map[string]interface{})["value"].(float64))
	glog.Infof("SearchAccountsByFiltersWithPaging: Full aggregations: %v\n", aggregations)
	glog.Infof("SearchAccountsByFiltersWithPaging: Total balance of matched accounts %v:", totalBalance)

	return accList, totalHits, totalBalance
}
