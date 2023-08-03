package repository

import (
	"chaobum-api/config"
	"context"
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"

	sq "github.com/Masterminds/squirrel"

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
	query := sq.Select("*").From("Photo")
	rows, err := query.RunWith(repo.db).Query()
	if err != nil {
		fmt.Printf("failed to run query. error: %s\n", err.Error())
		return nil, err
	}

	photos := []entity_port.PhotoPort{}
	for rows.Next() {
		photo := &entity.Photo{}
		if err := rows.Scan(&photo.ImageUrl, &photo.ShootingDate, &photo.CreatedAt, &photo.UpdatedAt); err != nil {
			fmt.Printf("failed to scan sql rows. error: %s\n", err.Error())
			break
		}
		photos = append(photos, *photo)
	}
	fmt.Printf("photos: %v\n", photos)

	return photos, nil
}

func (repo *PhotoRepository) SaveImageFile(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	defer file.Close()

	bucket, err := repo.storageClient.DefaultBucket()
	if err != nil {
		fmt.Printf("failed to get storage bucket. error: %s\n", err.Error())
		return err.Error(), err
	}

	fileUpload := bucket.Object(fileHeader.Filename)
	writer := fileUpload.NewWriter(repo.storageCtx)
	// writer.ObjectAttrs.ContentType = "application/json"
	// writer.ObjectAttrs.CacheControl = "no-cache"

	//firebase storageにファイルをアップロード (保存)
	_, err = io.Copy(writer, file)
	if err != nil {
		fmt.Printf("failed to io.Copy(). error: %s\n", err.Error())
		return err.Error(), err
	}

	if err := writer.Close(); err != nil {
		fmt.Printf("failed to close cloud storage writer. error: %s\n", err.Error())
		return err.Error(), err
	}

	imageUrl := "https://firebasestorage.googleapis.com/v0/b/" + config.Env.FIREBASE_PROJECT_ID + ".appspot.com/o/" + fileHeader.Filename + "?alt=media"

	return imageUrl, nil
}

func (repo *PhotoRepository) CreatePhoto(photo *entity.Photo) error {
	queryString, args, err := sq.Insert("Photo").Columns("imageUrl", "shootingDate", "createdAt", "updatedAt").Values(photo.ImageUrl, photo.ShootingDate, photo.CreatedAt, photo.UpdatedAt).ToSql()
	if err != nil {
		fmt.Printf("failed to create query. error: %s\n", err.Error())
		return err
	}

	_, err = repo.db.Exec(queryString, args...)
	if err != nil {
		fmt.Printf("failed to execute query. error: %s\n", err.Error())
		return err
	}

	return nil
}
