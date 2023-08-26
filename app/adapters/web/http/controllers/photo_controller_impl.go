package controller

import (
	view "chaobum-api/adapters/web/http/views"
	entity "chaobum-api/domains/entities"
	repository "chaobum-api/domains/repositories"
	usecase "chaobum-api/interactors/usecases/photo"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type photoControllerImpl struct {
	photoRepository          repository.PhotoRepository
	postPhotoService         *usecase.PostPhotoService
	updatePhotoService       *usecase.UpdatePhotoService
	deletePhotoService       *usecase.DeletePhotoService
	downloadImageFileService *usecase.DownloadImageFileService
}

func NewPhotoControllerImpl(photoRepository repository.PhotoRepository, postPhotoService *usecase.PostPhotoService, updatePhotoService *usecase.UpdatePhotoService, deletePhotoService *usecase.DeletePhotoService, downloadImageFileService *usecase.DownloadImageFileService) *photoControllerImpl {
	return &photoControllerImpl{photoRepository, postPhotoService, updatePhotoService, deletePhotoService, downloadImageFileService}
}

func (controller *photoControllerImpl) GetAllPhoto() http.HandlerFunc {
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

func (controller *photoControllerImpl) GetById() http.HandlerFunc {
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

func (controller *photoControllerImpl) PostPhoto() http.HandlerFunc {
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

func (controller *photoControllerImpl) UpdatePhoto() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			log.Printf("failed to get param id.\n")
			http.Error(w, "failed to get param id.\n", http.StatusBadRequest)
			return
		}

		var input view.PhotoInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			log.Printf("failed to read input from request body at update photo. error: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := controller.updatePhotoService.Handle(id, input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	return handler
}

func (controller *photoControllerImpl) DeletePhoto() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, ok := vars["id"]
		if !ok {
			log.Printf("failed to get param id.\n")
			http.Error(w, "failed to get param id.\n", http.StatusBadRequest)
			return
		}

		err := controller.deletePhotoService.Handle(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}

	return handler
}

func (controller *photoControllerImpl) DownloadImageFile() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var input view.DownloadImageFileInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			log.Printf("failed to read input from request body at update photo. error: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fileName := strings.Split(strings.Split(input.ImageUrl, "/")[7], "?")[0]

		storageReader, err := controller.downloadImageFileService.Handle(fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer storageReader.Close()

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		if _, err := io.Copy(w, storageReader); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	return handler
}
