package structs

import (
	"time"

	"gorm.io/gorm"
)

// User - Стандартный пользователь приложения
// dsdasdsadasd
type User struct {
	gorm.Model
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email" gorm:"unique"`
	HashedPassword string    `json:"hashed_password"`
	BirthDate      time.Time `json:"birth_date"`
}
