package usecase

type IDeletePhotoService interface {
	Handle(id string) error
}
