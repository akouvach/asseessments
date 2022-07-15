package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {

	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
		Dsn:         app.config.db.dsn,
		Port:        app.config.port,
	}

	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)

}
