package packets

import "time"

type ListUsersResponse struct {
	Users []*UserDetails `json:"users"`
	Total int            `json:"total"`
}

type UserDetails struct {
	ID            string    `json:"id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	IsAdmin       bool      `json:"isAdmin"`
	IsLocked      bool      `json:"isLocked"`
	LoginAttempts int       `json:"loginAttempts"`
	CreatedAt     time.Time `json:"createdAt" gorm:"column:created_at"`
}

type TokenResponse struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
