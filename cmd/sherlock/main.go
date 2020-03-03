package main

import (
	"os"

	"bakku.org/sherlock/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := web.NewServer(port)
	server.Start()
}
