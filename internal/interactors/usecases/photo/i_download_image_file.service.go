package usecase

import "cloud.google.com/go/storage"

type IDownloadImageFileService interface {
	Handle(fileName string) (*storage.Reader, error)
}
