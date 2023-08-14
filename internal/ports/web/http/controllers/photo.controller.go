package port

import "net/http"

type PhotoController interface {
	GetAllPhoto() http.HandlerFunc
	GetById() http.HandlerFunc
	PostPhoto() http.HandlerFunc
}
