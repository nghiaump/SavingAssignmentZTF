package main

import (
	"github.com/elastic/go-elasticsearch/v7"
	"log"
	"os"
)

const UserPort = ":50051"
const ElasticSearchPort = ":9200"
const ContainerElasticSearchEnv = "CONTAINER_ES_HOST"

func main() {
	// Lấy giá trị của biến môi trường
	addressESContainer := os.Getenv(ContainerElasticSearchEnv)
	if addressESContainer == "" {
		log.Println("Biến môi trường CONTAINER_ES_HOST không được cung cấp.")
		return
	}

	esClient, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://" + addressESContainer + ElasticSearchPort},
	})

	if err != nil {
		log.Println("Error creating Elasticsearch client:", err)
		return
	} else {
		log.Println("Connect thanh cong toi ElasticSearch")
	}

	userServiceHandler := NewUserServiceHandler(esClient)
	StartUserServer(userServiceHandler, UserPort)
}
