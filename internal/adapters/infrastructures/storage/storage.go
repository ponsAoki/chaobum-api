package storage

import (
	"chaobum-api/config"
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

type StorageClient struct {
	Client *storage.Client
	Ctx    context.Context
}

func NewStorageClient() (*StorageClient, error) {
	ctx := context.Background()
	config := &firebase.Config{
		StorageBucket: config.Env.FIREBASE_PROJECT_ID + ".appspot.com",
	}
	opts := option.WithCredentialsFile("chaobum-api-local-firebase-adminsdk.json")
	app, err := firebase.NewApp(ctx, config, opts)
	if err != nil {
		log.Fatalf("failed to initialize firebase app. error: %s\n", err.Error())
	}

	client, err := app.Storage(ctx)
	if err != nil {
		log.Fatalf("failed to initialize storage app. error: %s\n", err.Error())
		return nil, err
	}

	return &StorageClient{Client: client, Ctx: ctx}, nil
}
