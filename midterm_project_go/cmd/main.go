package main

import (
	"log"
	_ "midterm_cloud_project_2024/config"
	serverPkg "midterm_cloud_project_2024/internal/server"
)

func main() {
	server, err := serverPkg.New()
	if err != nil {
		log.Fatal(err)
	}

	server.Run()
}
