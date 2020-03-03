package models

type User struct {
	Base
	Names    string `validate:"required" json:"names,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `validate:"required" json:"password,omitempty"`
	Email    string `validate:"required,email" json:"email,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Address  string `json:"address,omitempty"`
	Visits   int    `json:"visits,omitempty"`
	Location string `json:"location,omitempty"`
	RoleID   int    `gorm:"index" json:"roleId,omitempty"`
	Role     string
}
