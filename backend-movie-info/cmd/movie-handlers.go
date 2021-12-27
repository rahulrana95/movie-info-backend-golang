package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print(errors.New("Invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movie, err := app.models.DB.Get(id)

	if err != nil {
		app.logger.Print(errors.New("err"))
		app.errorJSON(w, err)
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")

}

func (app *application) getMovies(w http.ResponseWriter, r *http.Request) {

	movie, err := app.models.DB.All()

	if err != nil {
		app.logger.Print(errors.New("err"))
		app.errorJSON(w, err)
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {
}
func (app *application) addMovie(w http.ResponseWriter, r *http.Request) {
}
func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {
}
