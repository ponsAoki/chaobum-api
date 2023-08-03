package port

import "net/http"

type PhotoController interface {
	GetAllPhoto() http.HandlerFunc
	PostPhoto() http.HandlerFunc
}
