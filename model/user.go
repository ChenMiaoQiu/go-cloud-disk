package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string
	PasswordDigest string
	NickName       string
	Status         string
	Avatar         string `gorm:"size:1000"`
}

const (
	// PasswordCount password encryption difficulty
	PasswordCount = 12
	// admin User
	StatusAdmin = "common_admin"
	// active User
	StatusActiveUser = "active"
	// inactive User
	StatusInactiveUser = "inactive"
	// Suspend User
	StatusSuspendUser = "suspend"
)

// SetPassword encrypt user password to save data
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCount)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword check user password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}
