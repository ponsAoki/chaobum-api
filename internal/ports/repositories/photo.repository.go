package port

import (
	view "chaobum-api/internal/adapters/web/http/views"
	entity "chaobum-api/internal/domains/entities"
	"mime/multipart"
)

type PhotoRepositoryPort interface {
	FindAllPhoto() ([]entity.IPhoto, error)
	FindById(id string, user *entity.Photo) (entity.IPhoto, error)
	SaveImageFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	CreatePhoto(imageUrl, shootingDate string) error
	UpdatePhoto(id string, input view.PhotoInput) error
	DeletePhoto(id string) error
}
