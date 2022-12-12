package app

import (
	"blubooks/model"
	"blubooks/repository"
	"blubooks/util/validator"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/segmentio/ksuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func readBook(app *App, w http.ResponseWriter, r *http.Request, id string) (*model.Book, error) {

	book, err := repository.ReadBook(app.db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return nil, err
		}
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)

		return nil, err
	}

	return book, nil
}
func readSection(app *App, w http.ResponseWriter, r *http.Request) (*model.Section, error) {

	id := chi.URLParam(r, "id")

	section, err := repository.ReadSection(app.db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return nil, err
		}
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)

		return nil, err
	}

	return section, nil
}

func getSections(app *App, w http.ResponseWriter, r *http.Request) (error, model.SectionDtos) {
	clientID := chi.URLParam(r, "id")

	list, err := repository.ListSections(app.db, clientID)
	if err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)
		return err, nil
	}
	if list == nil {
		var list model.SectionDtos
		return nil, list
	}
	return nil, list.ToDto()
}

func getSectionsTitles(app *App, w http.ResponseWriter, r *http.Request, bookID string) (model.SectionNavis, error) {

	list, err := repository.ListSectionsTitlesRec(app.db, bookID)
	if err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)
		return nil, err
	}
	if list == nil {
		var list model.SectionNavis
		return list, nil
	}
	return list, err
}

func (app *App) PageReadBook(w http.ResponseWriter, r *http.Request) {

	bookID := chi.URLParam(r, "id")

	sections, err := getSectionsTitles(app, w, r, bookID)
	if err != nil {
		return
	}

	book, err := readBook(app, w, r, bookID)
	if err != nil {
		return
	}

	var page model.PageDto
	page.Content.Book = book.ToDto()
	page.Content.Sections = sections.ToSectionNaviDto()

	if err := json.NewEncoder(w).Encode(page); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataCreationFailure, err)
		return
	}

}

func (app *App) PageReadSection(w http.ResponseWriter, r *http.Request) {

	section, err := readSection(app, w, r)
	if err != nil {
		return
	}

	book, err := readBook(app, w, r, section.BookID)
	if err != nil {
		return
	}

	sections, err := getSectionsTitles(app, w, r, section.BookID)
	if err != nil {
		return
	}

	var page model.PageDto
	page.Content.Book = book.ToDto()
	page.Content.Section = section.ToMarkdownDto()
	page.Content.Sections = sections.ToSectionNaviDto()

	if err := json.NewEncoder(w).Encode(page); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataCreationFailure, err)
		return
	}

}

func (app *App) ReadSection(w http.ResponseWriter, r *http.Request) {

	section, err := readSection(app, w, r)
	if err != nil {
		return
	}

	if err := json.NewEncoder(w).Encode(section.ToDto()); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataCreationFailure, err)
		return
	}

}

func (app *App) CreateSection(w http.ResponseWriter, r *http.Request) {
	form := &model.SectionForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusUnprocessableEntity, appErrFormDecodingFailure, err)
		return
	}

	if err := app.validator.Struct(form); err != nil {
		log.Warn(err)

		resp := validator.ToErrResponse(err, nil)
		if resp == nil {
			printError(app, w, http.StatusInternalServerError, appErrFormErrResponseFailure, err)
			return
		}
		respBody, err := json.Marshal(resp)
		if err != nil {
			log.Warn(err)
			printError(app, w, http.StatusInternalServerError, appErrJsonCreationFailure, err)
			return
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(respBody)
		return
	}

	sectionModel, err := form.ToModel()
	if err != nil {
		log.Warn(err)
		printError(app, w, http.StatusUnprocessableEntity, appErrFormDecodingFailure, err)
		return
	}

	id := chi.URLParam(r, "id")

	book, _ := repository.ReadBook(app.db, id)
	if book != nil {
		sectionModel.BookID = book.ID
	} else {
		subSection, err := repository.ReadSection(app.db, id)
		if err != nil {
			printError(app, w, http.StatusInternalServerError, appErrDataAccessFailure, err)
			return
		}

		sectionModel.SectionID = subSection.ID
		sectionModel.BookID = subSection.BookID
	}

	sectionModel.ID = ksuid.New().String()

	section, err := repository.CreateSection(app.db, sectionModel)
	if err != nil {

		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataUpdateFailure, err)
		return
	}
	log.Infof("New book created: %d", section.ID)
	w.WriteHeader(http.StatusCreated)
}

func (app *App) UpdateSection(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	form := &model.SectionForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Warn(err)
		printError(app, w, http.StatusUnprocessableEntity, appErrFormDecodingFailure, err)
		return
	}

	if err := app.validator.Struct(form); err != nil {
		log.Warn(err)

		resp := validator.ToErrResponse(err, nil)
		if resp == nil {
			printError(app, w, http.StatusInternalServerError, appErrFormErrResponseFailure, err)
			return
		}
		respBody, err := json.Marshal(resp)
		if err != nil {
			log.Warn(err)
			printError(app, w, http.StatusInternalServerError, appErrJsonCreationFailure, err)
			return
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(respBody)
		return
	}

	sectionModel, err := form.ToModel()
	if err != nil {
		log.Warn(err)
		printError(app, w, http.StatusUnprocessableEntity, appErrFormDecodingFailure, err)
		return
	}
	sectionModel.ID = id
	if err := repository.UpdateSection(app.db, sectionModel); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Warn(err)
		printError(app, w, http.StatusInternalServerError, appErrDataUpdateFailure, err)
		return
	}
	log.Infof("Section updated: %d", id)
	w.WriteHeader(http.StatusAccepted)
}
