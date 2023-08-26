package usecase

import (
	view "chaobum-api/adapters/web/http/views"
	repository "chaobum-api/domains/repositories"
)

type UpdatePhotoService struct {
	photoRepo repository.PhotoRepository
}

func NewUpdatePhotoService(photoRepo repository.PhotoRepository) *UpdatePhotoService {
	return &UpdatePhotoService{photoRepo}
}

func (updatePhotoService *UpdatePhotoService) Handle(id string, input view.PhotoInput) error {
	err := updatePhotoService.photoRepo.UpdatePhoto(id, input)
	if err != nil {
		return err
	}

	return nil
}
