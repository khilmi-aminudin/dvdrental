package db

import (
	"database/sql"
	"dvdrental/helper"
	"fmt"
	"os"
	"time"
)

func Connect() *sql.DB {
	var (
		dbName     = os.Getenv("DB_NAME")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbDriver   = os.Getenv("DB_DRIVER")
	)
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbDriver, dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open(dbDriver, connStr)
	helper.LogErrorAndPanic(err)

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetConnMaxLifetime(time.Minute * 60)

	return db
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			helper.Logger().Error(errRollback.Error())
		}
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			helper.Logger().Error(errCommit.Error())
		}
	}
}
