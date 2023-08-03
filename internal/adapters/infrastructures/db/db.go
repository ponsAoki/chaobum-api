package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient struct {
	DB *sql.DB
}

func NewDBClient(driverName, dataSourceName string) (*DBClient, error) {
	//connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}

	//test db connection
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
