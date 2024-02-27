package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
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

func CreateIndexingRequest(req *pb.SavingAccount) esapi.IndexRequest {
	out, _ := uuid.NewUUID()
	req.Id = out.String()
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

	log.Printf("Test json marshal %v", string(jsonStr))
	indexReq := esapi.IndexRequest{
		Index:   ESSavingIndex,
		Body:    strings.NewReader(string(jsonStr)),
		Refresh: "true",
	}

	return indexReq
}
func GetAllAccountsByUserIDHelper(userID string, client *elasticsearch.Client) []*pb.SavingAccount {
	var r map[string]interface{}
	var buf bytes.Buffer
	// match -> gan giong
	// term -> chinh xac giong
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"user_id.keyword": userID,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
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
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	accList := []*pb.SavingAccount{}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		//log.Printf(" * ID=%s:\n", hit.(map[string]interface{})["_id"])
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		// Convert the bytes data to JSON
		jsonData, err := json.Marshal(doc)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
		}

		// Convert JSON to struct
		accStruct := pb.SavingAccount{}
		if err := json.Unmarshal(jsonData, &accStruct); err != nil {
			log.Println("Error unmarshalling document in response:", err)
		} else {
			log.Printf("after unmarshaled: %v", accStruct)
			accList = append(accList, &accStruct)
		}
	}

	log.Println(strings.Repeat("=", 37))
	return accList
}

func SearchAccountsByFiltersHelper(filterObj *pb.Filter, client *elasticsearch.Client) []*pb.SavingAccount {
	query := GenerateQuery(filterObj)
	var r map[string]interface{}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
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
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)

	accList := []*pb.SavingAccount{}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		//log.Printf(" * ID=%s:\n", hit.(map[string]interface{})["_id"])
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		// Convert the bytes data to JSON
		jsonData, err := json.Marshal(doc)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
		}

		// Convert JSON to struct
		accStruct := pb.SavingAccount{}
		if err := json.Unmarshal(jsonData, &accStruct); err != nil {
			log.Println("Error unmarshalling document in response:", err)
		} else {
			log.Printf("after unmarshaled: %v", accStruct)
			accList = append(accList, &accStruct)
		}
	}

	log.Println(strings.Repeat("=", 37))
	return accList
}
