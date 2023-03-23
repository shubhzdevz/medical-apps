package models

type Report struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	PatientID    uint   `json:"patient_id"`
	DoctorID     uint   `json:"doctor_id"`
	LabTechID    uint   `json:"lab_tech_id"`
	TestType     string `json:"test_type"`
	TestResult   string `json:"test_result"`
	TestDate     string `json:"test_date"`
	CreatedAt    string `json:"created_at"`
}

// TableName specifies the table name for the Report model
func (r *Report) TableName() string {
	return "reports"
}
