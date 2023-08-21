package http

import (
	"chaobum-api/adapters/infrastructures/db"
	"log"
	"net/http"
)

type HttpAdapter struct{}

func NewHttpAdapter() *HttpAdapter {
	return &HttpAdapter{}
}

func (httpAdapter *HttpAdapter) Run(dbClient *db.DBClient) {
	err := http.ListenAndServe(":8080", httpAdapter.InitRouter(dbClient))
	if err != nil {
		log.Fatalf("server failed to run on :8080.\nerror: %s", err.Error())
	}
}
