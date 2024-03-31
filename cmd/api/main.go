package main

import (
	"GoComputeFlow/pkg/api"
	"time"
)

func main() {
	api.StartServer()
	time.Sleep(time.Hour)
}
