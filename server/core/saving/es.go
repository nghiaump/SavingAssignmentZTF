package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	pb "github.com/nghiaump/SavingAssignmentZTF/protobuf"
	"log"
	"strings"
)

func GetAllAccountsByUserIDHelper(userID string, client *elasticsearch.Client) []map[string]interface{} {
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
	//// Print the ID and document source for each hit.
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	//}
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s:\n", hit.(map[string]interface{})["_id"])
		doc := hit.(map[string]interface{})["_source"].(map[string]interface{})

		log.Println("Try to parse to SavingAccount")
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
		}
	}

	log.Println(strings.Repeat("=", 37))
	return nil
}
