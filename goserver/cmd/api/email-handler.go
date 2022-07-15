package main

import (
	"net/http"
	"uscabe/utils"
)

func (app *application) SendMail(w http.ResponseWriter, r *http.Request) {

	utils.SendMail()

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	ok := jsonResp{
		OK: true,
	}
	err := app.writeJSON(w, http.StatusOK, ok, "usuario")
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
