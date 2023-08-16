package usecase

import view "chaobum-api/internal/adapters/web/http/views"

type IUpdatePhotoService interface {
	Handle(id string, input view.PhotoInput) error
}
