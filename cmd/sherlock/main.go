package main

import (
	"log"
	"os"

	"bakku.org/sherlock/web"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	templateDir := os.Getenv("TEMPLATE_DIR")
	if templateDir == "" {
		log.Fatalln("error: TEMPLATE_DIR not given")
		return
	}

	server := web.NewServer(port, templateDir)
	server.Start()
}
