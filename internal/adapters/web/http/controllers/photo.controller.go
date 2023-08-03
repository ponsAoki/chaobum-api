package controller

import (
	post_photo_port "chaobum-api/internal/ports/interactors/usecases/photo"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PhotoController struct {
	getAllPhotosService post_photo_port.GetAllPhotosService
	postPhotoService    post_photo_port.PostPhotoServicePort
}

func NewPhotoController(getAllPhotosService post_photo_port.GetAllPhotosService, postPhotoService post_photo_port.PostPhotoServicePort) *PhotoController {
	return &PhotoController{getAllPhotosService, postPhotoService}
}

func (controller *PhotoController) GetAllPhoto() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		photos, err := controller.getAllPhotosService.Handle()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(photos)
	}

	return handler
}

func (controller *PhotoController) PostPhoto() http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			fmt.Println("failed to read form file")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		inputShootingDate := r.FormValue("shootingDate")
		shootingDate, err := time.Parse("2006-01-02 15:04:00", inputShootingDate)
		if err != nil {
			fmt.Println("failed to parse shooting date")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		photo, err := controller.postPhotoService.Handle(file, fileHeader, shootingDate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(photo)
	}

	return handler
}
