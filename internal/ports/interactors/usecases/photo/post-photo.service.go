package port

import (
	"mime/multipart"
)

type PostPhotoServicePort interface {
	Handle(image multipart.File, fileHeader *multipart.FileHeader, shootingDate string) error
}
