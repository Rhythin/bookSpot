package entities

type User struct {
	Email    string `gorm:"not null"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`
}
