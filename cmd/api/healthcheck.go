package main

import (
	"net/http"
)

func (a *applicationDependencies) healthcheckHandler(w http.ResponseWriter,
	r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": a.config.environment,
		"version":     appVersion,
	}
	err := a.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
