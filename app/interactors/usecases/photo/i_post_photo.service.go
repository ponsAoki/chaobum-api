package usecase

import (
	"mime/multipart"
)

type IPostPhotoService interface {
	Handle(image multipart.File, fileHeader *multipart.FileHeader, shootingDate string) error
}
