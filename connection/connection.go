package connection

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goapp-service/entity"
	_ "goapp-service/entity"
	logger "goapp-service/logger"
	"time"
)

const (
	dbDriverName = "mysql"
	dbUserName   = "root"
	dbUserPass   = "root"
	dbURL        = "localhost:3306"
	dbSchema     = "creative"
)

func CreateDBConnection() *sql.DB {
	infoLogger := logger.InfoLogger
	errorLogger := logger.ErrorLogger
	conn, err := sql.Open(dbDriverName, prepareUrl(""))
	if err != nil {
		errorLogger.Printf("%v", err)
	}
	conn.SetMaxIdleConns(20)
	conn.SetConnMaxLifetime(20)
	conn.SetConnMaxLifetime(time.Minute * 5)
	infoLogger.Printf("connection established")
	return conn
}

func prepareUrl(dbInstance string) string {
	if dbInstance == "" {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUserName, dbUserPass, dbURL, dbSchema)
	} else {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUserName, dbUserPass, dbURL, dbSchema)
	}
}

func InsertToDB(db *sql.DB, query string, lc entity.AppInfo) (*entity.AppInfo, error) {
	infoLogger := logger.InfoLogger
	infoLogger.Printf("Entering into insert with query %v with params %v", query, lc)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	stmt, err := db.PrepareContext(ctx, query)
	errorLogger := logger.ErrorLogger
	if err != nil {
		errorLogger.Printf("Error with %s preparing SQL statement", err)
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			errorLogger.Printf("Error with %s closing stmt", err)
		}
	}(stmt)
	var data = entity.AppInfo{}
	res, err := stmt.ExecContext(ctx, lc.AppId, lc.AppName)
	if err != nil {
		errorLogger.Printf("Error %v when inserting the rows into the table")
		return nil, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		errorLogger.Printf("Error %v when finding the rows affected the table")
		return nil, err
	}
	infoLogger.Printf("%d rows created ", affected)
	return &data, nil
}

func ReadAppInfoByID(db *sql.DB, query string, appId string) (*entity.AppInfo, error) {
	infoLogger := logger.InfoLogger
	infoLogger.Printf("Entering into insert with query %v with params %v", query, appId)
	errorLogger := logger.ErrorLogger
	var data = entity.AppInfo{}
	res, err := db.Query(query, appId)
	if err != nil {
		errorLogger.Printf("Error %v when inserting the rows into the table")
		return nil, err
	}
	for res.Next() {
		err := res.Scan(&data.AppId, &data.AppName)
		if err != nil {
			errorLogger.Printf("Error %v when scanning thr response", err)
		}
	}
	err = db.Close()
	if err != nil {
		return nil, err
	}
	return &data, nil
}
