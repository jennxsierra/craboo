package main

import (
	"errors"
	"net/http"

	"github.com/jennxsierra/craboo/internal/data"
)

// Handler to list all appointments (GET) and add new appointment (POST)
func (a *applicationDependencies) appointmentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// Perform encoding of the appointments data
	case http.MethodGet:
		dataEnvelope := envelope{"appointments": data.Appointments}
		err := a.writeJSON(w, http.StatusOK, dataEnvelope, nil)
		if err != nil {
			a.serverErrorResponse(w, r, err)
		}
	
	// Perform decoding of incoming JSON data to create a new appointment
	case http.MethodPost:
		var input data.Appointment
		err := a.readJSON(w, r, &input)
		if err != nil {
			a.badRequestResponse(w, r, err)
			return
		}
		// Basic validation (Require PatientName and DoctorName)
		if input.PatientName == "" || input.DoctorName == "" {
			a.badRequestResponse(w, r, errors.New("Patient and Doctor names are required."))
			return
		}
		// Pretend to add to DB (append to static slice)
		input.ID = int64(len(data.Appointments) + 1)
		data.Appointments = append(data.Appointments, input)
		err = a.writeJSON(w, http.StatusCreated, envelope{"appointment": input}, nil)
		if err != nil {
			a.serverErrorResponse(w, r, err)
		}
	default:
		a.methodNotAllowedResponse(w, r)
	}
}
