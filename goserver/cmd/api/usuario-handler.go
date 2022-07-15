package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"uscabe/models"
	"uscabe/utils"

	"github.com/julienschmidt/httprouter"
)

func UsuarioConvertir(datos map[string]interface{}) models.Usuario {

	var usu models.Usuario
	var err error

	log.Println("Ya tengo los datos en UsuarioConvertir", datos)

	usu.UsuarioId, err = strconv.ParseInt(utils.ExtractDataFromJson(datos, "UsuarioId", usu), 10, 64)
	if err != nil {
		fmt.Println("Error de conversion UsuarioId", err)
	}
	usu.Nombre = utils.ExtractDataFromJson(datos, "Nombre", usu)
	usu.Apellido = utils.ExtractDataFromJson(datos, "Apellido", usu)
	usu.Email = utils.ExtractDataFromJson(datos, "Email", usu)
	usu.PaisId = utils.ExtractDataFromJson(datos, "PaisId", usu)
	usu.Pass = utils.ExtractDataFromJson(datos, "Pass", usu)

	// switch v := ide.(type) {
	// case int64:
	// 	fmt.Println("int64", v)
	// case int:
	// 	fmt.Println("int", v)
	// case string:
	// 	fmt.Println("string", v)
	// default:
	// 	fmt.Println("otro: ", reflect.TypeOf(ide))

	// }
	// if ok {
	// 	usu.UsuarioId = id
	// } else {
	// 	fmt.Println(ok)

	// }

	// usu.Email = ExtractDataFromJson(datos, "Email")
	// usu.Nombre = ExtractDataFromJson(datos, "Nombre")
	// usu.Apellido = ExtractDataFromJson(datos, "Apellido")
	// usu.Pass = ExtractDataFromJson(datos, "Pass")
	// usu.PaisId = ExtractDataFromJson(datos, "PaisId")

	fmt.Println("Datos del usuario: ", usu)

	return usu

}

func (app *application) UsuarioPut(w http.ResponseWriter, r *http.Request) {

	/* voy a cargar los datos recibidos en un map */
	var result map[string]interface{}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal([]byte(buf), &result)
	if err != nil {
		log.Println("error en el unmarshal de usuario", err)
	}

	log.Println("Ya tengo los datos recibidos por post", result)

	/* ahora tengo que ver de cómo convertir dichos datos (que vienen en string)
	al tipo adecuado según la estructura */

	/* le mando a una función los valores para que la convierta.
	esta función me va a devolver un usuario cargado en la estructura correspondiente */

	usu := UsuarioConvertir(result)

	fmt.Println(usu)

	// tipoGenerico := Translate(usu)
	// log.Println("traducido con reflection", tipoGenerico)

	// var usuario models.Usuario
	// err := json.NewDecoder(r.Body).Decode(&usuario)

	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }
	// log.Println(usuario)

	/* debo agregar un registro */

	usuario, err := app.models.DB.UsuarioPut(usu)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = app.writeJSON(w, http.StatusOK, usuario, "usuario")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// type jsonResp struct {
	// 	OK bool `json:"ok"`
	// }

	// ok := jsonResp{
	// 	OK: true,
	// }

	// err = app.writeJSON(w, http.StatusOK, ok, "response")
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }

}

func (app *application) UsuarioGetOne(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	usuarioId, err := strconv.ParseInt(params.ByName("usuarioId"), 10, 64)
	if err != nil {
		app.logger.Print(err.Error())
		return
	}
	if usuarioId < 0 {
		err := errors.New("parámetro de usuarioId inválido, debe ser mayor a cero")
		app.logger.Print(err.Error())
		app.errorJSON(w, err)
		return
	}
	app.logger.Println("usuarioId is", usuarioId)

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

	if usuarioId > 0 {

		usuario, err := app.models.DB.UsuarioGetOne(usuarioId)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = app.writeJSON(w, http.StatusOK, usuario, "usuario")
		if err != nil {
			app.errorJSON(w, err)
			return
		}

	}

}

func (app *application) UsuarioGetAll(w http.ResponseWriter, r *http.Request) {

	usuarios, err := app.models.DB.UsuarioGetAll()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, usuarios, "usuarios")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}

func (app *application) UsuarioDelete(w http.ResponseWriter, r *http.Request) {

	ps := r.Context().Value("params").(httprouter.Params)

	id, err := strconv.ParseInt(ps.ByName("usuarioId"), 10, 64)

	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// log.Println(usuario)

	/* debo agregar un registro */

	err = app.models.DB.UsuarioDelete(id)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	ok := jsonResp{
		OK: true,
	}
	err = app.writeJSON(w, http.StatusOK, ok, "usuario")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// err = app.writeJSON(w, http.StatusOK, ok, "response")
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }

}
