package main

import (
	"net/http"

	"collab/pkg/config"

	"collab/pkg/collab"
	"collab/pkg/http/rest"
	"collab/pkg/logger"
	"collab/pkg/storage"
)

const appVersion string = "0.0.0"

func main() {
	logger := logger.InitSugarLogger()
	defer logger.Sync() // Flushes buffer, if any
	logger.Info("Verloop Collab Service")
	logger.Info("Initializing application for connecting to database and setting up API routes.")

	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	s, err := storage.NewStorage(config.GetDatabaseURI(), config.GetDatabaseName())
	if err != nil {
		panic(err)
	}

	pl := collab.NewcollabService(s)
	app := rest.NewAppContext(pl)
	logger.Infow("verloop Collabs Service Ready", "version", appVersion)
	logger.Fatal(http.ListenAndServe(":80", app.GetRouter()))
}
