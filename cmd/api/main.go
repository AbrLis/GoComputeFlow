package main

import (
	"GoComputeFlow/pkg/api"
	"GoComputeFlow/pkg/calculator"
	"GoComputeFlow/pkg/database"
	"time"
)

func main() {
	_ = database.OpenDB()
	calculator.CreateCalculators()
	api.StartServer()
	time.Sleep(time.Hour)
}
