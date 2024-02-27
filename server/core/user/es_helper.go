package main

import (
	"encoding/json"
	"fmt"
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

	log.Printf("Test json marshal %v", string(jsonStr))
	indexReq := esapi.IndexRequest{
		Index:   indexName,
		Body:    strings.NewReader(string(jsonStr)),
		Refresh: "true",
	}

	return indexReq
}
