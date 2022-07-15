package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		//call next middleware with new context
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	// func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)

	router.HandlerFunc(http.MethodGet, "/v1/paises", app.PaisGetAll)

	router.HandlerFunc(http.MethodGet, "/v1/grupo/:grupoId", app.GrupoGetOne)
	router.HandlerFunc(http.MethodGet, "/v1/grupos", app.GrupoGetAll)

	// router.HandlerFunc(http.MethodDelete, "/v1/usuario/:usuarioId", app.UsuarioDelete)
	// http.Handle("/v1/usuario/:usuarioId", secure.ThenFunc(app.UsuarioDelete))
	router.DELETE("/v1/usuarios/:usuarioId", app.wrap(secure.ThenFunc(app.UsuarioDelete)))
	// router.HandlerFunc(http.MethodPut, "/v1/usuario/:usuarioId", app.UsuarioPut)
	router.PUT("/v1/usuarios/:usuarioId", app.wrap(secure.ThenFunc(app.UsuarioPut)))
	router.HandlerFunc(http.MethodGet, "/v1/usuarios/:usuarioId", app.UsuarioGetOne)
	router.HandlerFunc(http.MethodPost, "/v1/usuarios/0", app.UsuarioPut)
	// router.HandlerFunc(http.MethodGet, "/v1/usuarios", app.UsuarioGetAll)
	router.GET("/v1/usuarios", app.wrap(secure.ThenFunc(app.UsuarioGetAll)))

	router.HandlerFunc(http.MethodGet, "/v1/sendmail", app.SendMail)

	return app.enableCORS(router)

}
