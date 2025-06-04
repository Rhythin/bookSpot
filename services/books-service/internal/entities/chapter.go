package entities

import "github.com/rhythin/bookspot/services/shared/custommodel"

type Chapter struct {
	custommodel.CustomModel
	Title  string `gorm:"not null"`
	BookID string `gorm:"not null"`
	Number int    `gorm:"not null"`

	Book Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}
