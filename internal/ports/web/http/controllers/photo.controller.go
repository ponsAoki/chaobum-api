package port

import "net/http"

type PhotoController interface {
	GetAllPhoto() http.HandlerFunc
	GetById() http.HandlerFunc
	PostPhoto() http.HandlerFunc
	UpdatePhoto() http.HandlerFunc
	DeletePhoto() http.HandlerFunc
}
