package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type env struct {
	FIREBASE_PROJECT_ID   string
	FIREBASE_PRIVATE_KEY  string
	FIREBASE_CLIENT_EMAIL string
	DB_DRIVER             string
	DATA_SOURCE_NAME      string
}

var Env env

func init() {
	godotenv.Load(".env")

	Env.FIREBASE_PROJECT_ID = os.Getenv("FIREBASE_PROJECT_ID")
	Env.FIREBASE_PRIVATE_KEY = strings.ReplaceAll(os.Getenv("FIREBASE_PRIVATE_KEY"), "\\n", "\n")
	Env.FIREBASE_CLIENT_EMAIL = os.Getenv("FIREBASE_CLIENT_EMAIL")
	Env.DB_DRIVER = os.Getenv("DB_DRIVER")
	Env.DATA_SOURCE_NAME = os.Getenv("DATA_SOURCE_NAME")
}
