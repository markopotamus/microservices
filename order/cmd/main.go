package main

import (
	"log"

	"github.com/markopotamus/microservices/order/config"
	"github.com/markopotamus/microservices/order/internal/adapters/db"
	"github.com/markopotamus/microservices/order/internal/adapters/grpc"
	"github.com/markopotamus/microservices/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSrcUrl())
	if err != nil {
		log.Fatalf("failed to connect to database, err=%v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetAppPort())
	grpcAdapter.Run()
}
