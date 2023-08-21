package port

import (
	"chaobum-api/adapters/infrastructures/db"
	"net/http"
)

type HttpPort interface {
	Run(dbClient *db.DBClient)
	InitRouter(dbClient *db.DBClient) http.Handler
}
