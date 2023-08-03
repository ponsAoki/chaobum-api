package port

import (
	entity "chaobum-api/internal/adapters/domains/entities"
	"mime/multipart"
	"time"
)

type PostPhotoServicePort interface {
	Handle(image multipart.File, fileHeader *multipart.FileHeader, shootingDate time.Time) (*entity.Photo, error)
}
