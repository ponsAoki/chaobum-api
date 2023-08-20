package usecase

import (
	repository_port "chaobum-api/internal/ports/repositories"

	"cloud.google.com/go/storage"
)

type DownloadImageFileService struct {
	photoRepo repository_port.PhotoRepositoryPort
}

func NewDownloadImageFileService(photoRepo repository_port.PhotoRepositoryPort) *DownloadImageFileService {
	return &DownloadImageFileService{photoRepo}
}

func (downloadImageFileService *DownloadImageFileService) Handle(fileName string) (*storage.Reader, error) {
	storageReader, err := downloadImageFileService.photoRepo.DownloadImageFile(fileName)
	if err != nil {
		return nil, err
	}

	return storageReader, nil
}
