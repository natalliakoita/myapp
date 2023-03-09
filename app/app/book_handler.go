package app

import (
	"fmt"
	"myapp/model"
	"myapp/repository"
	"net/http"

	"github.com/jinzhu/gorm"
)

const (
	appErrDataAccessFailure      = "data access failure"
	appErrJsonCreationFailure    = "json creation failure"
	appErrDataCreationFailure    = "data creation failure"
	appErrFormDecodingFailure    = "form decoding failure"
	appErrDataUpdateFailure      = "data update failure"
	appErrFormErrResponseFailure = "form error response failure"
	appErrUintRequestFailure     = "id request failure"
)

func (a *App) HandleListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := repository.ListBooks(a.db)
	if err != nil {
		RespondError(w, a, err, http.StatusInternalServerError, appErrDataAccessFailure)
		return
	}

	if books == nil {
		fmt.Fprint(w, "[]")
		return
	}

	dtos := books.ToDto()
	RespondJSON(w, r, a, &dtos, http.StatusInternalServerError, appErrJsonCreationFailure)
}

func (a *App) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	form := model.BookForm{}
	ParseRequestBody(w, r, a, &form, http.StatusUnprocessableEntity, appErrFormDecodingFailure)

	bookModel, err := form.ToModel()
	if err != nil {
		RespondError(w, a, err, http.StatusUnprocessableEntity, appErrFormDecodingFailure)
		return
	}

	book, err := repository.CreateBook(a.db, bookModel)
	if err != nil {
		RespondError(w, a, err, http.StatusInternalServerError, appErrDataCreationFailure)
		return
	}

	a.logger.Info().Msgf("New book created: %d", book.ID)
	w.WriteHeader(http.StatusCreated)
}

func (a *App) HandleReadBook(w http.ResponseWriter, r *http.Request) {
	id, err := ParseUint(w, r, a)
	if err != nil {
		RespondError(w, a, err, http.StatusUnprocessableEntity, appErrUintRequestFailure)
	}

	book, err := repository.ReadBook(a.db, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		RespondError(w, a, err, http.StatusInternalServerError, appErrDataAccessFailure)
		return
	}

	dto := book.ToDto()
	RespondJSON(w, r, a, &dto, http.StatusInternalServerError, appErrJsonCreationFailure)
}

func (a *App) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := ParseUint(w, r, a)
	if err != nil {
		RespondError(w, a, err, http.StatusUnprocessableEntity, appErrUintRequestFailure)
	}

	form := &model.BookForm{}
	ParseRequestBody(w, r, a, &form, http.StatusUnprocessableEntity, appErrFormDecodingFailure)

	bookModel, err := form.ToModel()
	if err != nil {
		RespondError(w, a, err, http.StatusUnprocessableEntity, appErrFormDecodingFailure)
	}

	bookModel.ID = id
	if err := repository.UpdateBook(a.db, bookModel); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		RespondError(w, a, err, http.StatusInternalServerError, appErrDataUpdateFailure)
		return
	}

	a.logger.Info().Msgf("Book updated: %d", id)
	w.WriteHeader(http.StatusAccepted)
}

func (a *App) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := ParseUint(w, r, a)
	if err != nil {
		RespondError(w, a, err, http.StatusUnprocessableEntity, appErrUintRequestFailure)
	}

	if err := repository.DeleteBook(a.db, id); err != nil {
		RespondError(w, a, err, http.StatusInternalServerError, appErrDataAccessFailure)
		return
	}

	a.logger.Info().Msgf("Book deleted: %d", id)
	w.WriteHeader(http.StatusAccepted)
}
