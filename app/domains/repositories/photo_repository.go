package repository

import (
	view "chaobum-api/adapters/web/http/views"
	entity "chaobum-api/domains/entities"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type PhotoRepository interface {
	FindAllPhoto() ([]entity.Photo, error)
	FindById(id string, user *entity.Photo) (*entity.Photo, error)
	SaveImageFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	CreatePhoto(imageUrl, shootingDate string) error
	UpdatePhoto(id string, input view.PhotoInput) error
	DeletePhoto(id string) error
	DownloadImageFile(fileName string) (*storage.Reader, error)
}
