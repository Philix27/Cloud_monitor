package user

import (
	"time"

	"gorm.io/gorm"
)



var DB *gorm.DB
var err error

const DNS = ""


type User struct {
	gorm.Model 
	ID        uint64    `gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	first_name string `json:"firstname"`
	last_name string `json:"firstname"`
	email string	`json:"firstname"`
}
