package main

import (
	"log"
	serverPkg "midterm_cloud_project_2024/internal/server"
)

func main() {
	server, err := serverPkg.New()
	if err != nil {
		log.Fatal(err)
	}

	server.Run()
}
