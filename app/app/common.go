package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func ParseRequestBody(w http.ResponseWriter, r *http.Request, a *App, v interface{}, statusCode int) {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		RespondError(w, a, fmt.Errorf("parse request body: %w", err), statusCode)
		return
	}
}

func RespondJSON(w http.ResponseWriter, r *http.Request, a *App, v interface{}, statusCode int) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		RespondError(w, a, fmt.Errorf("data encode: %w", err), statusCode)
		return
	}
}

func RespondError(w http.ResponseWriter, a *App, err error, statusCode int) {
	a.logger.Warn().Err(err).Msg("")

	w.WriteHeader(statusCode)
	fmt.Fprintf(w, `{"error": "%v"}`, err)
}

func ParseUint(w http.ResponseWriter, r *http.Request, a *App) (uint, error) {
	q1 := chi.URLParam(r, "id")
	_ = q1
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)
	if err != nil || id == 0 {
		a.logger.Info().Msgf("can not parse ID: %v", id)

		w.WriteHeader(http.StatusUnprocessableEntity)
		return uint(id), err
	}

	return uint(id), nil
}
