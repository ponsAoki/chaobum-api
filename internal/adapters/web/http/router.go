package http

import (
	"chaobum-api/internal/adapters/infrastructures/db"
	"chaobum-api/internal/adapters/infrastructures/storage"
	controller "chaobum-api/internal/adapters/web/http/controllers"
	middleware "chaobum-api/internal/adapters/web/http/middlewares"
	usecase "chaobum-api/internal/interactors/usecases/photo"
	repository_port "chaobum-api/internal/ports/repositories"
	controller_port "chaobum-api/internal/ports/web/http/controllers"
	repository "chaobum-api/internal/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (httpAdapter *HttpAdapter) InitRouter(dbClient *db.DBClient) http.Handler {
	storageClient, err := storage.NewStorageClient()
	if err != nil {
		log.Fatalf("failed to initialize firebase storage client.\nerror: %s", err.Error())
	}
	var photoRepository repository_port.PhotoRepositoryPort = repository.NewPhotoRepository(*storageClient.Client, storageClient.Ctx, dbClient.DB)
	var postPhotoService usecase.IPostPhotoService = usecase.NewPostPhotoService(photoRepository)
	var updatePhotoService usecase.IUpdatePhotoService = usecase.NewUpdatePhotoService(photoRepository)
	var deletePhotoService usecase.IDeletePhotoService = usecase.NewDeletePhotoService(photoRepository)
	var downloadImageFileService usecase.IDownloadImageFileService = usecase.NewDownloadImageFileService(photoRepository)
	var photoController controller_port.PhotoController = controller.NewPhotoController(photoRepository, postPhotoService, updatePhotoService, deletePhotoService, downloadImageFileService)

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
