package usecase

import (
	repository_port "chaobum-api/ports/repositories"
)

type DeletePhotoService struct {
	photoRepo repository_port.PhotoRepositoryPort
}

func NewDeletePhotoService(photoRepo repository_port.PhotoRepositoryPort) *DeletePhotoService {
	return &DeletePhotoService{photoRepo}
}

func (deletePhotoService *DeletePhotoService) Handle(id string) error {
	err := deletePhotoService.photoRepo.DeletePhoto(id)
	if err != nil {
		return err
	}

	return nil
}
