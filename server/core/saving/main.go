package main

import (
	"flag"
	"github.com/golang/glog"
	"os"
)

const SavingPort = ":50052"
const ElasticSearchPort = ":9200"
const ContainerElasticSearchEnv = "CONTAINER_ES_HOST"

func main() {
	// Init glog
	os.Args = append(os.Args, "-logtostderr=true")
	flag.Parse()
	defer glog.Flush()
	
	esClient, _ := CreateESClient()
	InitIndex(ESSavingIndex, esClient)
	db := CreateMySQLClient()
	defer db.Close()
	savingServiceHandler := NewSavingServiceHandler(esClient, db)
	StartSavingServer(savingServiceHandler, SavingPort)
}
