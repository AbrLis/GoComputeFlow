package main

import (
	"log"

	"GoComputeFlow/internal/api"
	"GoComputeFlow/internal/api/apiConfig"
	"GoComputeFlow/internal/calculator"
	"GoComputeFlow/internal/database"
	"GoComputeFlow/internal/frontend"
	"GoComputeFlow/internal/worker"
	workerServer "GoComputeFlow/internal/worker/server"
)

func main() {
	// Открытие базы данных
	if err := database.OpenDB(); err != nil {
		log.Fatal("Error opening database: ", err)
	}

	// Установка таймаутов с прошлого запуска и создание воркеров
	timeouts := database.GetTimeouts()
	worker.CreateWorker(timeouts.AddTimeout, timeouts.SubtractTimeout, timeouts.MutiplyTimeout, timeouts.DivideTimeout)

	// Запуск grpc сервера воркеров
	workerServer.StartGRPCServerWorker(calculator.GRPChost, calculator.GRPCport)

	// Создание калькуляторов - для работы необходима открытая бд и сервер воркеров
	calculator.CreateCalculators()

	// Запуск API
	go api.StartServer(apiConfig.HostPath, apiConfig.PortHost)

	// Запуск фронта
	frontend.StartFront()
}
