package usecase

import (
	entity_port "chaobum-api/internal/ports/domains/entities"
	repository_port "chaobum-api/internal/ports/repositories"
)

type GetAllPhotosService struct {
	photoRepo repository_port.PhotoRepositoryPort
}

func NewGetAllPhotosService(photoRepo repository_port.PhotoRepositoryPort) *GetAllPhotosService {
	return &GetAllPhotosService{photoRepo}
}

func (service *GetAllPhotosService) Handle() ([]entity_port.PhotoPort, error) {
	photos, err := service.photoRepo.FindAllPhoto()
	if err != nil {
		return nil, err
	}

	return photos, nil
}
