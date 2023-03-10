package app

import (
	"fmt"
	"myapp/model"
	"net/http"

	"github.com/jinzhu/gorm"
)

func (a *App) HandleListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := a.svcBook.GetListBook(r.Context())
	if err != nil {
		RespondError(w, a, fmt.Errorf("data access failure: %w", err), http.StatusInternalServerError)
		return
	}

	RespondJSON(w, r, a, &books, http.StatusInternalServerError)
}

func (a *App) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	bookForm := model.BookForm{}
	ParseRequestBody(w, r, a, &bookForm, http.StatusUnprocessableEntity)

	book, err := a.svcBook.CreateBook(r.Context(), &bookForm)
	if err != nil {
		RespondError(w, a, fmt.Errorf("data creation failure: %w", err), http.StatusInternalServerError)
		return
	}

	a.logger.Info().Msgf("New book created: %d", book.ID)
	w.WriteHeader(http.StatusCreated)
}

func (a *App) HandleReadBook(w http.ResponseWriter, r *http.Request) {
	id, err := ParseUint(w, r, a)
	if err != nil {
		RespondError(w, a, fmt.Errorf("id request failure: %w", err), http.StatusUnprocessableEntity)
	}

	book, err := a.svcBook.GetBookByID(r.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		RespondError(w, a, fmt.Errorf("data access failure: %w", err), http.StatusInternalServerError)
		return
	}

	RespondJSON(w, r, a, &book, http.StatusInternalServerError)
}

func (a *App) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := ParseUint(w, r, a)
	if err != nil {
		RespondError(w, a, fmt.Errorf("id request failure: %w", err), http.StatusUnprocessableEntity)
	}

	bookForm := &model.BookForm{}
	ParseRequestBody(w, r, a, &bookForm, http.StatusUnprocessableEntity)

	if err := a.svcBook.UpdateBook(r.Context(), id, bookForm); err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		RespondError(w, a, fmt.Errorf("data update failure: %w", err), http.StatusInternalServerError)
		return
	}

	a.logger.Info().Msgf("Book updated: %d", id)
	w.WriteHeader(http.StatusAccepted)
}

func (a *App) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := ParseUint(w, r, a)
	if err != nil {
		RespondError(w, a, fmt.Errorf("id request failure: %w", err), http.StatusUnprocessableEntity)
	}

	if err := a.svcBook.DeleteBook(r.Context(), id); err != nil {
		RespondError(w, a, fmt.Errorf("data access failure: %w", err), http.StatusInternalServerError)
		return
	}

	a.logger.Info().Msgf("Book deleted: %d", id)
	w.WriteHeader(http.StatusAccepted)
}
