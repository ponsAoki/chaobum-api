package http

import (
	"chaobum-api/internal/adapters/infrastructures/db"
	"chaobum-api/internal/adapters/infrastructures/storage"
	usecase "chaobum-api/internal/adapters/interactors/usecases/photo"
	repository "chaobum-api/internal/adapters/repositories"
	controller "chaobum-api/internal/adapters/web/http/controllers"
	usecase_port "chaobum-api/internal/ports/interactors/usecases/photo"
	repository_port "chaobum-api/internal/ports/repositories"
	controller_port "chaobum-api/internal/ports/web/http/controllers"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (httpAdapter *HttpAdapter) InitRouter(dbClient *db.DBClient) http.Handler {
	storageClient, err := storage.NewStorageClient()
	if err != nil {
		log.Fatalf("failed to initialize firebase storage client.\nerror: %s", err.Error())
	}
	var photoRepository repository_port.PhotoRepositoryPort = repository.NewPhotoRepository(*storageClient.Client, storageClient.Ctx, dbClient.DB)
	var getAllPhotosService usecase_port.GetAllPhotosService = usecase.NewGetAllPhotosService(photoRepository)
	var postPhotoService usecase_port.PostPhotoServicePort = usecase.NewPostPhotoService(photoRepository)
	var photoController controller_port.PhotoController = controller.NewPhotoController(getAllPhotosService, postPhotoService)

	r := chi.NewRouter()

	r.Get("/photo/get-all", photoController.GetAllPhoto())
	r.Post("/photo/post-photo", photoController.PostPhoto())

	return r
}
