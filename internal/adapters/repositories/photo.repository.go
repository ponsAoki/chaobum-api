package repository

import (
	"chaobum-api/config"
	"context"
	"database/sql"
	"io"
	"log"
	"mime/multipart"
	"time"

	entity "chaobum-api/internal/adapters/domains/entities"
	entity_port "chaobum-api/internal/ports/domains/entities"

	firebase_storage "firebase.google.com/go/storage"
)

type PhotoRepository struct {
	storageClient firebase_storage.Client
	storageCtx    context.Context
	db            *sql.DB
}

func NewPhotoRepository(storageClient firebase_storage.Client, storageCtx context.Context, db *sql.DB) *PhotoRepository {
	return &PhotoRepository{storageClient, storageCtx, db}
}

func (repo *PhotoRepository) FindAllPhoto() ([]entity_port.PhotoPort, error) {
	rows, err := repo.db.Query("SELECT * FROM Photo")
	if err != nil {
		log.Printf("failed to run query. error: %s\n", err.Error())
		return nil, err
	}

	photos := []entity_port.PhotoPort{}
	for rows.Next() {
		photo := &entity.Photo{}
		if err := rows.Scan(&photo.ID, &photo.ImageUrl, &photo.ShootingDate, &photo.CreatedAt, &photo.UpdatedAt); err != nil {
			log.Printf("failed to scan sql rows. error: %s\n", err.Error())
			break
		}
		photos = append(photos, *photo)
	}

	return photos, nil
}

func (repo *PhotoRepository) FindById(id string, photo *entity.Photo) (entity_port.PhotoPort, error) {
	err := repo.db.QueryRow("SELECT * FROM Photo WHERE id = ?", id).Scan(&photo.ID, &photo.ImageUrl, &photo.ShootingDate, &photo.CreatedAt, &photo.UpdatedAt)
	if err != nil {
		log.Printf("failed to scan sql row in FindById. error: %s", err.Error())
		return nil, err
	}

	return photo, nil
}

func (repo *PhotoRepository) SaveImageFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	defer file.Close()

	bucket, err := repo.storageClient.DefaultBucket()
	if err != nil {
		log.Printf("failed to get storage bucket. error: %s\n", err.Error())
		return err.Error(), err
	}

	fileUpload := bucket.Object(fileHeader.Filename)
	writer := fileUpload.NewWriter(repo.storageCtx)
	// writer.ObjectAttrs.ContentType = "application/json"
	// writer.ObjectAttrs.CacheControl = "no-cache"

	//firebase storageにファイルをアップロード (保存)
	_, err = io.Copy(writer, file)
	if err != nil {
		log.Printf("failed to io.Copy(). error: %s\n", err.Error())
		return err.Error(), err
	}

	if err := writer.Close(); err != nil {
		log.Printf("failed to close cloud storage writer. error: %s\n", err.Error())
		return err.Error(), err
	}

	imageUrl := "https://firebasestorage.googleapis.com/v0/b/" + config.Env.FIREBASE_PROJECT_ID + ".appspot.com/o/" + fileHeader.Filename + "?alt=media"

	return imageUrl, nil
}

func (repo *PhotoRepository) CreatePhoto(imageUrl, shootingDate string) error {
	_, err := repo.db.Exec("INSERT INTO Photo (imageUrl, shootingDate, createdAt, updatedAt) VALUES (?, ?, ?, ?)", imageUrl, shootingDate, time.Now(), time.Now())
	if err != nil {
		log.Printf("failed to execute query. error: %s\n", err.Error())
		return err
	}

	return nil
}
