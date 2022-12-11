package app

import (
	"blubooks/model"
	"blubooks/repository"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func getCollectionsFromClient(app *App, w http.ResponseWriter, r *http.Request) (error, model.CollectionDtos) {
	clientID := chi.URLParam(r, "id")

	list, err := repository.ListCollectionByClientID(app.db, clientID)
	if err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)
		return err, nil
	}
	if list == nil {
		var list model.CollectionDtos
		return nil, list
	}
	return nil, list.ToDto()
}

func readCollection(app *App, w http.ResponseWriter, r *http.Request) (error, *model.CollectionDto) {

	id := chi.URLParam(r, "id")

	collection, err := repository.ReadCollection(app.db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return err, nil
		}
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)

		return err, nil
	}

	return nil, collection.ToDto()
}
func getBooks(app *App, w http.ResponseWriter, r *http.Request) (error, model.BookDtos) {
	clientID := chi.URLParam(r, "id")

	list, err := repository.ListBooks(app.db, clientID)
	if err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)
		return err, nil
	}
	if list == nil {
		var list model.BookDtos
		return nil, list
	}
	return nil, list.ToDto()
}

func (app *App) GetCollectionsFromClient(w http.ResponseWriter, r *http.Request) {

	err, list := getCollectionsFromClient(app, w, r)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(list); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataCreationFailure, err)
		return
	}

}

func (app *App) PageGetCollectionsFromClient(w http.ResponseWriter, r *http.Request) {

	err, list := getCollectionsFromClient(app, w, r)
	if err != nil {
		return
	}

	err, client := readClient(app, w, r)
	if err != nil {
		return
	}

	var page model.PageDto
	page.Content.Collections = list
	page.Content.Client = client

	if err := json.NewEncoder(w).Encode(page); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataCreationFailure, err)
		return
	}

}

func (app *App) PageGetCollection(w http.ResponseWriter, r *http.Request) {

	err, books := getBooks(app, w, r)
	if err != nil {
		return
	}

	err, collection := readCollection(app, w, r)
	if err != nil {
		return
	}

	var page model.PageDto
	page.Content.Collection = collection
	page.Content.Books = books

	if err := json.NewEncoder(w).Encode(page); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataCreationFailure, err)
		return
	}

}
