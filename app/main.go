package main

import (
	"chaobum-api/adapters/infrastructures/db"
	"chaobum-api/adapters/web/http"
	"chaobum-api/config"
	"log"
)

func main() {
	dbClient, err := db.NewDBClient(config.Env.DB_DRIVER)
	if err != nil {
		log.Fatalf("failed to initiate database connection: %v", err)
	}
	defer dbClient.CloseDbConnection()

	var httpAdapter http.HttpPort = http.NewHttpAdapter()
	httpAdapter.Run(dbClient)
}
