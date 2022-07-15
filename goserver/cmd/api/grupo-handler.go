package main

import (
	"errors"
	"net/http"
	"strconv"
	"uscabe/models"

	"github.com/julienschmidt/httprouter"
)

func (app *application) GrupoGetOne(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	grupoId, err := strconv.ParseInt(params.ByName("grupoId"), 10, 64)
	if err != nil {
		app.logger.Print(errors.New("parámetro de grupoId inválido"))
		app.errorJSON(w, err)
		return
	}
	app.logger.Println("grupoId is", grupoId)

	// audit := models.Auditoria{
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// 	UsuCreate: "andres@litim.com",
	// 	UsuUpdate: "andres@litim.com",
	// }

	// ramasUtilizadas := models.Ramas{
	// 	Lobatos: false,
	// 	Scouts:  false,
	// 	Raiders: true,
	// 	Rovers:  true,
	// }

	grupo := models.Grupo{
		GrupoId: grupoId,
		Nombre:  "Nro 23",
		// Anio:            2022,
		// Audit:           audit,
		// RamasUtilizadas: ramasUtilizadas,
	}

	err = app.writeJSON(w, http.StatusOK, grupo, "grupo")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) GrupoGetAll(w http.ResponseWriter, r *http.Request) {

	grupos, err := app.models.DB.GrupoGetAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, grupos, "grupos")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
