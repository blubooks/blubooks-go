package app

import (
	"blubooks/model"
	"blubooks/repository"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

func listClients(app *App, w http.ResponseWriter, r *http.Request) (error, model.ClientDtos) {

	list, err := repository.ListClients(app.db)
	if err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)
		return err, nil
	}
	if list == nil {
		var list model.ClientDtos
		return nil, list
	}
	return nil, list.ToDto()
}
func readClient(app *App, w http.ResponseWriter, r *http.Request) (error, *model.ClientDto) {

	id := chi.URLParam(r, "id")

	client, err := repository.ReadClient(app.db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return err, nil
		}
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)

		return err, nil
	}

	return nil, client.ToDto()
}

func (app *App) ListClients(w http.ResponseWriter, r *http.Request) {

	err, list := listClients(app, w, r)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(list); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataCreationFailure, err)
		return
	}
}
func (app *App) PageListClients(w http.ResponseWriter, r *http.Request) {
	err, list := listClients(app, w, r)
	if err != nil {
		return
	}

	var page model.PageDto
	page.Content.Clients = list

	if err := json.NewEncoder(w).Encode(page); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataCreationFailure, err)
		return
	}
}

func (app *App) ReadClient(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	client, err := repository.ReadClient(app.db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)

		return
	}
	dto := client.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		log.Warn(err)

		printError(app, w, http.StatusInternalServerError, appErrJsonCreationFailure, err)
		return
	}
}

func (app *App) PageReadClient(w http.ResponseWriter, r *http.Request) {
	err, client := readClient(app, w, r)
	if err != nil {
		return
	}
	err, collections := getCollectionsFromClient(app, w, r)
	if err != nil {
		return
	}

	var page model.PageDto
	page.Content.Client = client
	page.Content.Collections = collections

	if err := json.NewEncoder(w).Encode(page); err != nil {
		log.Warn(err)

		printError(app, w, http.StatusInternalServerError, appErrJsonCreationFailure, err)
		return
	}

}
