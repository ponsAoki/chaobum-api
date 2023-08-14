package http

import (
	"chaobum-api/internal/adapters/infrastructures/db"
	"chaobum-api/internal/adapters/infrastructures/storage"
	usecase "chaobum-api/internal/adapters/interactors/usecases/photo"
	repository "chaobum-api/internal/adapters/repositories"
	controller "chaobum-api/internal/adapters/web/http/controllers"
	middleware "chaobum-api/internal/adapters/web/http/middlewares"
	usecase_port "chaobum-api/internal/ports/interactors/usecases/photo"
	repository_port "chaobum-api/internal/ports/repositories"
	controller_port "chaobum-api/internal/ports/web/http/controllers"
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
	var postPhotoService usecase_port.PostPhotoServicePort = usecase.NewPostPhotoService(photoRepository)
	var photoController controller_port.PhotoController = controller.NewPhotoController(photoRepository, postPhotoService)

	r := mux.NewRouter()

	r.Handle("/photo/get-all", middleware.SetCors(photoController.GetAllPhoto()))
	r.Handle("/photo/get-by-id/{id}", middleware.SetCors(photoController.GetById()))
	r.Handle("/photo/post-photo", middleware.SetCors(photoController.PostPhoto()))

	http.Handle("/", r)

	return r
}
