package models

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// TableName specifies the table name for the User model
func (u *User) TableName() string {
	return "users"
}
