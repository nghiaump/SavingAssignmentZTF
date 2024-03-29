package main

import (
	"flag"
	"github.com/golang/glog"
	"os"
)

const UserPort = ":50051"
const ElasticSearchPort = ":9200"
const ContainerElasticSearchEnv = "CONTAINER_ES_HOST"

func main() {
	// Init glog
	os.Args = append(os.Args, "-logtostderr=true")
	flag.Parse()
	defer glog.Flush()

	esClient, _ := CreateESClient()
	InitIndex(ESUserIndex, esClient)
	db := CreateMySQLClient()
	defer db.Close()
	userServiceHandler := NewUserServiceHandler(esClient, db)
	StartUserServer(userServiceHandler, UserPort)
}
