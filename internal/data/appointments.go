package data

import "time"

// Appointment represents a medical appointment
type Appointment struct {
	ID          int64     `json:"id"`
	PatientName string    `json:"patient_name"`
	DoctorName  string    `json:"doctor_name"`
	Date        time.Time `json:"date"`
	Reason      string    `json:"reason"`
}

// Static pretend data for practice
var Appointments = []Appointment{
	{
		ID:          1,
		PatientName: "Alice Smith",
		DoctorName:  "Dr. John Doe",
		Date:        time.Date(2026, 2, 20, 10, 0, 0, 0, time.UTC),
		Reason:      "Annual Checkup",
	},
	{
		ID:          2,
		PatientName: "Bob Johnson",
		DoctorName:  "Dr. Jane Roe",
		Date:        time.Date(2026, 2, 21, 14, 30, 0, 0, time.UTC),
		Reason:      "Consultation",
	},
}
