package main

import (
	"github.com/Terrero89/runners_go/config"
	"github.com/Terrero89/runners_go/server"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	log.Println("Starting runners app...")

	log.Println("Initializing config...")
	config := config.InitConfig("runners")

	log.Println("Initilizing Database...")
	dbHandler := server.InitDatabase(config)

	log.Println("Starting server...")

	httpServer := server.InitHttpServer(config, dbHandler)

	httpServer.Start()
}
