package main

const SavingPort = ":50052"
const ElasticSearchPort = ":9200"
const ContainerElasticSearchEnv = "CONTAINER_ES_HOST"

func main() {
	esClient, _ := CreateESClient()
	InitIndex(ESSavingIndex, esClient)
	savingServiceHandler := NewSavingServiceHandler(esClient)
	StartSavingServer(savingServiceHandler, SavingPort)
}
