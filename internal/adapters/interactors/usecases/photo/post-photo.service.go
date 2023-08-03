package usecase

import (
	entity "chaobum-api/internal/adapters/domains/entities"
	repository_port "chaobum-api/internal/ports/repositories"
	"fmt"
	"mime/multipart"
	"time"
)

type PostPhotoService struct {
	photoRepo repository_port.PhotoRepositoryPort
}

func NewPostPhotoService(photoRepo repository_port.PhotoRepositoryPort) *PostPhotoService {
	return &PostPhotoService{photoRepo}
}

func (postPhotoService *PostPhotoService) Handle(file multipart.File, fileHeader *multipart.FileHeader, shootingDate time.Time) (*entity.Photo, error) {
	fmt.Printf("imageFile: %v\nshootingDate: %v\n", file, shootingDate)

	//ストレージにファイルをアップロード
	imageUrl, err := postPhotoService.photoRepo.SaveImageFile(file, fileHeader)
	if err != nil {
		return nil, err
	}

	newPhoto := entity.NewPhoto(imageUrl, shootingDate, time.Now(), time.Now())

	//新規photoデータの永続化
	err = postPhotoService.photoRepo.CreatePhoto(newPhoto)
	if err != nil {
		return nil, err
	}

	return newPhoto, nil
}
