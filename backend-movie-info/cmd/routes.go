package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	log.Fatal(http.ListenAndServe(":5000", router))

	return router
}
