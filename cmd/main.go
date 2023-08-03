package main

import (
	"chaobum-api/config"
	"chaobum-api/internal/adapters/infrastructures/db"
	"chaobum-api/internal/adapters/web/http"
	http_port "chaobum-api/internal/ports/web/http"
	"log"
)

func main() {
	dbClient, err := db.NewDBClient(config.Env.DB_DRIVER, config.Env.DATA_SOURCE_NAME)
	if err != nil {
		log.Fatalf("failed to initiate database connection: %v", err)
	}
	defer dbClient.CloseDbConnection()

	var httpAdapter http_port.HttpPort = http.NewHttpAdapter()
	httpAdapter.Run(dbClient)
}
