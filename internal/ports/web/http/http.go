package port

import (
	"chaobum-api/internal/adapters/infrastructures/db"
	"net/http"
)

type HttpPort interface {
	Run(dbClient *db.DBClient)
	InitRouter(dbClient *db.DBClient) http.Handler
}
