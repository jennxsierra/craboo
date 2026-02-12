package main

import (
	"fmt"
	"net/http"
)

func (a *applicationDependencies) healthcheckHandler(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", a.config.environment)
	fmt.Fprintf(w, "version: %s\n", appVersion)
}
