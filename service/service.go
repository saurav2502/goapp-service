package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goapp-service/connection"
	"goapp-service/entity"
	log "goapp-service/logger"
	"net/http"
)

func HandleGetAppInfoById(writer http.ResponseWriter, req *http.Request) {
	appId := mux.Vars(req)["appId"]
	infoLogger := log.InfoLogger
	infoLogger.Printf("Entering inside HandleGetAppInfoById with args %v", appId)
	infos := entity.CreateAppInfo()
	for _, app := range infos {
		if app.AppId == appId {
			log.DebugLogger.Printf("Request URI %v and Res %v", log.GetRequestURI(req), app)
			json.NewEncoder(writer).Encode(app)
		}
	}
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	infoLogger := log.InfoLogger
	appId := mux.Vars(r)["appId"]
	infoLogger.Printf("Entering inside HandleGet with args %v", appId)
	dbConnection := connection.CreateDBConnection()
	query := "select id as appId, name as appName from creative where id = ?"
	data, err := connection.ReadAppInfoByID(dbConnection, query, appId)
	if err != nil {
		log.ErrorLogger.Printf("Query failed: %v", err)
		return
	}
	infoLogger.Printf("Returning from HandleGet with response %v", *data)
	debugLogger := log.DebugLogger
	debugLogger.Printf("API Req %v and Res %v", log.GetRequestURI(r), *data)
	json.NewEncoder(w).Encode(data)
}
