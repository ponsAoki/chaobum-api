package usecase

import (
	repository "chaobum-api/domains/repositories"
)

type DeletePhotoService struct {
	photoRepo repository.PhotoRepository
}

func NewDeletePhotoService(photoRepo repository.PhotoRepository) *DeletePhotoService {
	return &DeletePhotoService{photoRepo}
}

func (deletePhotoService *DeletePhotoService) Handle(id string) error {
	err := deletePhotoService.photoRepo.DeletePhoto(id)
	if err != nil {
		return err
	}

	return nil
}
