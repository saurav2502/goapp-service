package controller

import (
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"goapp-service/config"
	log "goapp-service/logger"
	"goapp-service/service"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
}

func (a *App) Run() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./resources")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration config.Configurations
	errorLogger := log.ErrorLogger
	if err := viper.ReadInConfig(); err != nil {
		errorLogger.Printf("Error reading config file, %s", err)
	}
	viper.SetDefault("database.dbname", "test_db")
	err := viper.Unmarshal(&configuration)
	if err != nil {
		errorLogger.Printf("Unable to decode into struct, %v", err)
	}
	port := ":" + strconv.Itoa(configuration.Server.Port)
	log := log.InfoLogger
	log.Print("App is running on the port ", port)
	log.Fatal(http.ListenAndServe(port, a.Router))
}

func (app *App) RestHandlerController() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/app/{appId}", service.HandleGetAppInfoById)
	app.Router.HandleFunc("/app/v1/{appId}", service.HandleGet)
}
