package main

import (
	"GoComputeFlow/pkg/api"
	"GoComputeFlow/pkg/database"
	"time"
)

func main() {
	_ = database.OpenDB()
	api.StartServer()
	time.Sleep(time.Hour)
}
