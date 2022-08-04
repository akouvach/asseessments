package main

import (
	"net/http"
)

func (app *application) PaisGetAll(w http.ResponseWriter, r *http.Request) {

	valores, err := app.models.DB.PaisGetAll(app.config.db.name)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, valores, "paises")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}
