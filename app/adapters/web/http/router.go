package http

import (
	"chaobum-api/adapters/infrastructures/db"
	"chaobum-api/adapters/infrastructures/storage"
	controller "chaobum-api/adapters/web/http/controllers"
	middleware "chaobum-api/adapters/web/http/middlewares"
	repository "chaobum-api/domains/repositories"
	usecase "chaobum-api/interactors/usecases/photo"
	repository_impl "chaobum-api/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (httpAdapter *HttpAdapter) InitRouter(dbClient *db.DBClient) http.Handler {
	storageClient, err := storage.NewStorageClient()
	if err != nil {
		log.Fatalf("failed to initialize firebase storage client.\nerror: %s", err.Error())
	}
	var photoRepository repository.PhotoRepository = repository_impl.NewPhotoRepositoryImpl(*storageClient.Client, storageClient.Ctx, dbClient.DB)
	postPhotoService := usecase.NewPostPhotoService(photoRepository)
	updatePhotoService := usecase.NewUpdatePhotoService(photoRepository)
	deletePhotoService := usecase.NewDeletePhotoService(photoRepository)
	downloadImageFileService := usecase.NewDownloadImageFileService(photoRepository)
	var photoController controller.PhotoController = controller.NewPhotoControllerImpl(photoRepository, postPhotoService, updatePhotoService, deletePhotoService, downloadImageFileService)

	r := mux.NewRouter()

	r.Handle("/photo", middleware.SetCors(photoController.GetAllPhoto())).Methods(http.MethodGet)
	r.Handle("/photo/{id}", middleware.SetCors(photoController.GetById())).Methods(http.MethodGet)
	r.Handle("/photo", middleware.SetCors(photoController.PostPhoto())).Methods(http.MethodPost)
	r.Handle("/photo/{id}", middleware.SetCors(photoController.UpdatePhoto())).Methods(http.MethodOptions, http.MethodPut)
	r.Handle("/photo/{id}", middleware.SetCors(photoController.DeletePhoto())).Methods(http.MethodDelete)
	r.Handle("/photo/image_file_download", middleware.SetCors(photoController.DownloadImageFile())).Methods(http.MethodPost)

	http.Handle("/", r)

	return r
}
