package main

import (
	"GoComputeFlow/pkg/database"
	"fmt"
)

func main() {
	_ = database.OpenDB()
	id, _ := database.CreateNewUser("admin", "admin")
	fmt.Println("New user ID:", id)
}
