package structs

import (
	"time"

	"gorm.io/gorm"
)

// User - Стандартный пользователь приложения
type User struct {
	gorm.Model
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email" gorm:"unique"`
	HashedPassword string    `json:"hashed_password"`
	BirthDate      time.Time `json:"birth_date"`
	LastLoginDate  time.Time `json:"last_login_date"`
}

type UserLoginModel struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterModel struct {
	FirstName string    `json:"first_name" validate:"required"`
	LastName  string    `json:"last_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	BirthDate time.Time `json:"birth_date" validate:"required"`
}
