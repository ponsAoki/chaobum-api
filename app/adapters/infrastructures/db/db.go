package db

import (
	"chaobum-api/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	DB *sql.DB
}

func NewDBClient(driverName string) (*DBClient, error) {
	//connect
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", config.Env.DB_USER, config.Env.DB_PASSWORD, config.Env.DB_HOST, config.Env.DB_PORT, config.Env.DB_NAME)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failed: %v", err)
	}

	return &DBClient{DB: db}, nil
}

func (da *DBClient) CloseDbConnection() {
	err := da.DB.Close()
	if err != nil {
		log.Fatalf("db close failed: %v", err)
	}
}
