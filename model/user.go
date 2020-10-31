package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a User schema
type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:VARCHAR(191)" gorm:"unique"`
	Username string `json:"username" gorm:"type:VARCHAR(100)"gorm:"unique"`
	Password string `json:"password" gorm:"type:VARCHAR(191)" gorm:"->:false;<-:create"`
	Token    string `json:"token"`
}

type UserDto struct {
	Email    string `json:"email" gorm:"type:VARCHAR(191)" gorm:"unique"`
	Username string `json:"username" gorm:"type:VARCHAR(100)"gorm:"unique"`
	Token    string `json:"token"`
}

func makePassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash := makePassword(u.Password)
		u.Password = hash
	}
	return
}

// Claims represent the structure of the JWT token
type Claims struct {
	jwt.StandardClaims
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique"`
}
