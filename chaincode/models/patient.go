package models

type Patient struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// TableName specifies the table name for the Patient model
func (p *Patient) TableName() string {
	return "patients"
}
