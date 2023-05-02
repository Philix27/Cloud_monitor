package user

import (
	"time"

	"gorm.io/gorm"
)



type User struct {
	gorm.Model 
	ID        uint64    `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	First_name 	string 	`json:"first_name"`
	Last_name 	string 	`json:"last_name"`
	Email 		string		`json:"email"`
}
