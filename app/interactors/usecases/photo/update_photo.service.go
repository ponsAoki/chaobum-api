package usecase

import (
	view "chaobum-api/adapters/web/http/views"
	repository_port "chaobum-api/ports/repositories"
)

type UpdatePhotoService struct {
	photoRepo repository_port.PhotoRepositoryPort
}

func NewUpdatePhotoService(photoRepo repository_port.PhotoRepositoryPort) *UpdatePhotoService {
	return &UpdatePhotoService{photoRepo}
}

func (updatePhotoService *UpdatePhotoService) Handle(id string, input view.PhotoInput) error {
	err := updatePhotoService.photoRepo.UpdatePhoto(id, input)
	if err != nil {
		return err
	}

	return nil
}
