package app

import (
	"blubooks/repository"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func (app *App) GetCollections(w http.ResponseWriter, r *http.Request) {

	clientID := chi.URLParam(r, "id")
	if len(clientID) != 27 {
		printError(app, w, http.StatusInternalServerError, "Wrong ID", nil)
		return
	}
	list, err := repository.ListCollection(app.db, clientID)
	if err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)
		return
	}
	if list == nil {
		fmt.Fprint(w, "[]")
		return
	}
	dtos := list.ToDto()
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		printError(app, w, http.StatusInternalServerError, appErrJsonCreationFailure, err)

		return
	}
}
