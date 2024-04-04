package main

import (
	"GoComputeFlow/pkg/api"
	"GoComputeFlow/pkg/calculator"
	"GoComputeFlow/pkg/database"
	worker "GoComputeFlow/pkg/worker/server"
	"log"
	"time"
)

func main() {
	if err := database.OpenDB(); err != nil {
		log.Fatal("Error opening database: ", err)
	}

	calculator.CreateCalculators()
	api.StartServer(api.HostPath, api.PortHost)

	if err := worker.StartGRPCServerWorker("localhost", ":3001"); err != nil {
		log.Fatal("Error start server gRPC worker: ", err)
	}
	time.Sleep(time.Hour)
}
