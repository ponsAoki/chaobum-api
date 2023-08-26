package usecase

import (
	repository "chaobum-api/domains/repositories"

	"cloud.google.com/go/storage"
)

type DownloadImageFileService struct {
	photoRepo repository.PhotoRepository
}

func NewDownloadImageFileService(photoRepo repository.PhotoRepository) *DownloadImageFileService {
	return &DownloadImageFileService{photoRepo}
}

func (downloadImageFileService *DownloadImageFileService) Handle(fileName string) (*storage.Reader, error) {
	storageReader, err := downloadImageFileService.photoRepo.DownloadImageFile(fileName)
	if err != nil {
		return nil, err
	}

	return storageReader, nil
}
