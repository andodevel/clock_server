package models

import "time"

// User ...
type User struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `json:"name"`
	Username  string     `json:"username"`
}
