package main

import (
	"log"

	"github.com/maximis3d/issue-tracking-system/cmd/api"
)

func main() {

	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal()
	}
}
