package app

import (
	"blubooks/model"
	"blubooks/repository"
	"blubooks/server/service"
	"blubooks/util/tools"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	form := &model.UserLoginForm{}
	if app.checkForm(form, w, r) {
		return
	}

	user, err := repository.GetUserByEmail(app.db, form.Email)
	if err != nil {
		printError(app, w, http.StatusInternalServerError, "user & password not matched", err)
		return
	}
	if !tools.CheckPasswordHash(form.Password, user.Password) {
		printError(app, w, http.StatusInternalServerError, "user & password not matched", nil)
		return
	}

	jwt := service.Jwt{}
	token, err := jwt.CreateToken(*user)
	if err != nil {
		printError(app, w, http.StatusInternalServerError, "unable to create access token", err)
		return
	}

	dtos := token.ToDto(user)
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrJsonCreationFailure, err)
	}

}

func (app *App) RefreshLoginToken(w http.ResponseWriter, r *http.Request) {

	token := model.Token{}

	log.Printf("body: %d", r.Context())
	if err := json.NewDecoder(r.Body).Decode(&token); err != nil {
		printError(app, w, http.StatusUnprocessableEntity, appErrFormDecodingFailure, err)
		return
	}

	jwt := service.Jwt{}
	user, err := jwt.ValidateRefreshToken(token)
	if err != nil {
		printError(app, w, http.StatusNotAcceptable, "invalid token", err)
		return
	}

	token, err = jwt.CreateToken(user)
	if err != nil {
		printError(app, w, http.StatusInternalServerError, "unable to create access token", err)
		return
	}

	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error.message": "%v"}`, appErrJsonCreationFailure)
		return
	}
}
