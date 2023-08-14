package usecase

import (
	repository_port "chaobum-api/internal/ports/repositories"
	"mime/multipart"
)

type PostPhotoService struct {
	photoRepo repository_port.PhotoRepositoryPort
}

func NewPostPhotoService(photoRepo repository_port.PhotoRepositoryPort) *PostPhotoService {
	return &PostPhotoService{photoRepo}
}

func (postPhotoService *PostPhotoService) Handle(file multipart.File, fileHeader *multipart.FileHeader, shootingDate string) error {
	//ストレージにファイルをアップロード
	imageUrl, err := postPhotoService.photoRepo.SaveImageFile(file, fileHeader)
	if err != nil {
		return err
	}

	//新規photoデータの永続化
	err = postPhotoService.photoRepo.CreatePhoto(imageUrl, shootingDate)
	if err != nil {
		return err
	}

	return nil
}
