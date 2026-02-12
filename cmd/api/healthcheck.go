package main

import (
	"fmt"
	"net/http"
)

func (a *applicationDependencies) healthcheckHandler(w http.ResponseWriter,
	r *http.Request) {
	jsResponse := `{"status": "available", "environment": %q, "version": %q}`
	jsResponse = fmt.Sprintf(jsResponse, a.config.environment, appVersion)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsResponse))
}
