package app

import (
	"blubooks/repository"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (app *App) HandleListClients(w http.ResponseWriter, r *http.Request) {
	list, err := repository.ListClients(app.db)
	if err != nil {
		log.Warn(err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error.message": "%v"}`, appErrDataAccessFailure)
		return
	}
	if list == nil {
		fmt.Fprint(w, "[]")
		return
	}
	dtos := list.ToDto()
	if err := json.NewEncoder(w).Encode(dtos); err != nil {
		log.Warn(err)

		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error.message": "%v"}`, appErrJsonCreationFailure)
		return
	}
}
