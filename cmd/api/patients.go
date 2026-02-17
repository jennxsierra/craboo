package main

import (
	"errors"
	"net/http"

	"github.com/jennxsierra/lab3-json-handlers/internal/data"
)

// Handler to list all patients (GET) and add new patient (POST)
func (a *applicationDependencies) patientsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// Perform encoding of the patients data
	case http.MethodGet:
		dataEnvelope := envelope{"patients": data.Patients}
		err := a.writeJSON(w, http.StatusOK, dataEnvelope, nil)
		if err != nil {
			a.serverErrorResponse(w, r, err)
		}

	// Perform decoding of incoming JSON data to create a new patient
	case http.MethodPost:
		var input data.Patient
		err := a.readJSON(w, r, &input)
		if err != nil {
			a.badRequestResponse(w, r, err)
			return
		}
		// Basic validation (Require Name)
		if input.Name == "" {
			a.badRequestResponse(w, r, errors.New("Patient's name is required."))
			return
		}
		// Pretend to add to DB (append to static slice)
		input.ID = int64(len(data.Patients) + 1)
		data.Patients = append(data.Patients, input)
		err = a.writeJSON(w, http.StatusCreated, envelope{"patient": input}, nil)
		if err != nil {
			a.serverErrorResponse(w, r, err)
		}
	default:
		a.methodNotAllowedResponse(w, r)
	}
}
