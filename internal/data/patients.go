package data

// Patient represents a patient in the system
type Patient struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Static pretend data for practice
var Patients = []Patient{
	{ID: 1, Name: "Alice Smith", Age: 30},
	{ID: 2, Name: "Bob Johnson", Age: 45},
}
