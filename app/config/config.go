package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type env struct {
	FIREBASE_PROJECT_ID   string
	FIREBASE_PRIVATE_KEY  string
	FIREBASE_CLIENT_EMAIL string
	DB_DRIVER             string
	DB_USER               string
	DB_PASSWORD           string
	DB_HOST               string
	DB_PORT               int
	DB_NAME               string
	FRONT_URL             string
	DATA_SOURCE_NAME      string
}

var Env env

func init() {
	Env.FIREBASE_PROJECT_ID = os.Getenv("FIREBASE_PROJECT_ID")
	Env.FIREBASE_PRIVATE_KEY = strings.ReplaceAll(os.Getenv("FIREBASE_PRIVATE_KEY"), "\\n", "\n")
	Env.FIREBASE_CLIENT_EMAIL = os.Getenv("FIREBASE_CLIENT_EMAIL")
	Env.DB_DRIVER = os.Getenv("DB_DRIVER")
	Env.DB_USER = os.Getenv("DB_USER")
	Env.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	Env.DB_HOST = os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("failed to convert string DB PORT to int. error: %s", err.Error())
	}
	Env.DB_PORT = dbPort
	Env.DB_NAME = os.Getenv("DB_NAME")
	Env.FRONT_URL = os.Getenv("FRONT_URL")
}
