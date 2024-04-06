package main

import (
	"log"
	"time"

	"GoComputeFlow/pkg/api"
	"GoComputeFlow/pkg/calculator"
	"GoComputeFlow/pkg/database"
	"GoComputeFlow/pkg/worker"
	workerServer "GoComputeFlow/pkg/worker/server"
)

func main() {
	if err := database.OpenDB(); err != nil {
		log.Fatal("Error opening database: ", err)
	}

	workerServer.StartGRPCServerWorker(calculator.GRPChost, calculator.GRPCport)
	calculator.CreateCalculators()
	worker.CreateWorker()

	api.StartServer(api.HostPath, api.PortHost)
	time.Sleep(time.Hour)
}
