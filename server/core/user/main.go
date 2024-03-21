package main

const UserPort = ":50051"
const ElasticSearchPort = ":9200"
const ContainerElasticSearchEnv = "CONTAINER_ES_HOST"

func main() {
	esClient, _ := CreateESClient()
	InitIndex(ESUserIndex, esClient)
	db := CreateMySQLClient()
	defer db.Close()
	userServiceHandler := NewUserServiceHandler(esClient, db)
	StartUserServer(userServiceHandler, UserPort)
}
