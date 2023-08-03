package port

import (
	entity "chaobum-api/internal/adapters/domains/entities"
	entity_port "chaobum-api/internal/ports/domains/entities"
	"mime/multipart"
)

type PhotoRepositoryPort interface {
	FindAllPhoto() ([]entity_port.PhotoPort, error)
	SaveImageFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	CreatePhoto(photo *entity.Photo) error
}
