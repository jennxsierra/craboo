package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *applicationDependencies) routes() http.Handler {

	// setup a new router
	router := httprouter.New()

	// handle 404
	router.NotFound = http.HandlerFunc(a.notFoundResponse)

	// handle 405
	router.MethodNotAllowed = http.HandlerFunc(a.methodNotAllowedResponse)

	// setup routes
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", a.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/comments", a.createCommentHandler)

	// New APIs for appointments and patients (GET and POST)
	router.HandlerFunc(http.MethodGet, "/v1/appointments", a.appointmentsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/appointments", a.appointmentsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/patients", a.patientsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/patients", a.patientsHandler)

	return a.recoverPanic(router)
}
