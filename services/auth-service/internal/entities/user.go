package entities

import (
	"github.com/rhythin/bookspot/services/shared/custommodel"
)

type User struct {
	custommodel.CustomModel
	Username      string `gorm:"uniqueIndex;not null" json:"username"`
	Email         string `gorm:"not null" json:"email"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Password      string `gorm:"not null" json:"-"`
	Salt          string `gorm:"not null" json:"-"`
	LoginAttempts int    `gorm:"default:0" json:"login_attempts"`
	IsLocked      bool   `gorm:"default:false" json:"is_locked"`
	IsAdmin       bool   `gorm:"default:false" json:"is_admin"`
	TempToken     string `gorm:"type:text" json:"-"`
}
