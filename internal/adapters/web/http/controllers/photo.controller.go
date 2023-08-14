package controller

import (
	entity "chaobum-api/internal/adapters/domains/entities"
	post_photo_port "chaobum-api/internal/ports/interactors/usecases/photo"
	repository_port "chaobum-api/internal/ports/repositories"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PhotoController struct {
	photoRepository  repository_port.PhotoRepositoryPort
	postPhotoService post_photo_port.PostPhotoServicePort
}

func NewPhotoController(photoRepository repository_port.PhotoRepositoryPort, postPhotoService post_photo_port.PostPhotoServicePort) *PhotoController {
	return &PhotoController{photoRepository, postPhotoService}
}

func (controller *PhotoController) GetAllPhoto() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		photos, err := controller.photoRepository.FindAllPhoto()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(photos)
	}

	return handler
}

func (controller *PhotoController) GetById() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			log.Printf("failed to get param id.\n")
			http.Error(w, "failed to get param id.\n", http.StatusBadRequest)
			return
		}

		photo := &entity.Photo{}
		resPhoto, err := controller.photoRepository.FindById(id, photo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(resPhoto)
	}

	return handler
}

func (controller *PhotoController) PostPhoto() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.Printf("failed to read form file. error: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		shootingDate := r.FormValue("shootingDate")

		err = controller.postPhotoService.Handle(file, fileHeader, shootingDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}

	return handler
}
