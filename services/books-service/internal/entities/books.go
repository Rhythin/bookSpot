package entities

import "github.com/rhythin/bookspot/services/shared/custommodel"

type Book struct {
	custommodel.CustomModel
	Title       string    `gorm:"not null"`
	Author      string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Chapters    []Chapter `gorm:"foreignKey:BookID;references:ID"`
}
