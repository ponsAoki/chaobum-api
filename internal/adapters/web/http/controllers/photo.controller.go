package controller

import (
	view "chaobum-api/internal/adapters/web/http/views"
	entity "chaobum-api/internal/domains/entities"
	usecase "chaobum-api/internal/interactors/usecases/photo"
	repository_port "chaobum-api/internal/ports/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type PhotoController struct {
	photoRepository          repository_port.PhotoRepositoryPort
	postPhotoService         usecase.IPostPhotoService
	updatePhotoService       usecase.IUpdatePhotoService
	deletePhotoService       usecase.IDeletePhotoService
	downloadImageFileService usecase.IDownloadImageFileService
}

func NewPhotoController(photoRepository repository_port.PhotoRepositoryPort, postPhotoService usecase.IPostPhotoService, updatePhotoService usecase.IUpdatePhotoService, deletePhotoService usecase.IDeletePhotoService, downloadImageFileService usecase.IDownloadImageFileService) *PhotoController {
	return &PhotoController{photoRepository, postPhotoService, updatePhotoService, deletePhotoService, downloadImageFileService}
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

func (controller *PhotoController) UpdatePhoto() http.HandlerFunc {
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

func (controller *PhotoController) DeletePhoto() http.HandlerFunc {
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

func (controller *PhotoController) DownloadImageFile() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var input view.DownloadImageFileInput

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			log.Printf("failed to read input from request body at update photo. error: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fileName := strings.Split(strings.Split(input.ImageUrl, "/")[7], "?")[0]
		log.Printf("fileName: %s", fileName)

		storageReader, err := controller.downloadImageFileService.Handle(fileName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer storageReader.Close()

		// w.Header().Set("Content-Disposition", "attachment; filename=yourfilename.txt")
		log.Println(fmt.Sprintf("attachment; filename=%s", fileName))
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		// w.Header().Set("Content-Type", "text/plain")
		if _, err := io.Copy(w, storageReader); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	return handler
}
