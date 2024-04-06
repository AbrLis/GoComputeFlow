package main

import (
	"log"
	"time"

	"GoComputeFlow/internal/api"
	"GoComputeFlow/internal/calculator"
	"GoComputeFlow/internal/database"
	"GoComputeFlow/internal/worker"
	workerServer "GoComputeFlow/internal/worker/server"
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
