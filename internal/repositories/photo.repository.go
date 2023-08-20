package repository

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"time"

	"chaobum-api/config"
	view "chaobum-api/internal/adapters/web/http/views"
	entity "chaobum-api/internal/domains/entities"

	"cloud.google.com/go/storage"
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

func (repo *PhotoRepository) FindAllPhoto() ([]entity.IPhoto, error) {
	rows, err := repo.db.Query("SELECT * FROM Photo")
	if err != nil {
		log.Printf("failed to run query. error: %s\n", err.Error())
		return nil, err
	}

	photos := []entity.IPhoto{}
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

func (repo *PhotoRepository) FindById(id string, photo *entity.Photo) (entity.IPhoto, error) {
	err := repo.db.QueryRow("SELECT * FROM Photo WHERE id = ?", id).Scan(&photo.ID, &photo.ImageUrl, &photo.ShootingDate, &photo.CreatedAt, &photo.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("no photo records found. error: %s\n", err.Error())
		return nil, err
	}
	if err != nil {
		log.Printf("failed to scan sql row in FindById. error: %s\n", err.Error())
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
	_, err := repo.db.Exec("INSERT INTO photo (imageUrl, shootingDate, createdAt, updatedAt) VALUES (?, ?, ?, ?)", imageUrl, shootingDate, time.Now(), time.Now())
	if err != nil {
		log.Printf("failed to execute query to create photo. error: %s\n", err.Error())
		return err
	}

	return nil
}

func (repo *PhotoRepository) UpdatePhoto(id string, input view.PhotoInput) error {
	_, err := repo.db.Exec("UPDATE photo SET shootingDate = ? WHERE id = ?", input.ShootingDate, id)
	if err != nil {
		log.Printf("failed to execute query to update photo. error: %s\n", err.Error())
		return err
	}

	return nil
}

func (repo *PhotoRepository) DeletePhoto(id string) error {
	_, err := repo.db.Exec("DELETE FROM photo WHERE id = ?", id)
	if err != nil {
		log.Printf("failed to execute query to delete photo. error: %s\n", err.Error())
		return err
	}

	return nil
}

func (repo *PhotoRepository) DownloadImageFile(fileName string) (*storage.Reader, error) {
	bucket, err := repo.storageClient.DefaultBucket()
	if err != nil {
		log.Printf("failed to get storage bucket. error: %s\n", err.Error())
		return nil, err
	}

	storageReader, err := bucket.Object(fileName).NewReader(repo.storageCtx)
	if err != nil {
		log.Printf("failed to create storage reader at download image file. error: %s\n", err.Error())
		return nil, err
	}

	return storageReader, nil
}
