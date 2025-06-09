package entities

import "github.com/rhythin/bookspot/services/shared/custommodel"

type Chapter struct {
	custommodel.CustomModel
	Title  string `gorm:"not null" json:"title"`
	BookID string `gorm:"not null" json:"-"`
	Number int    `gorm:"not null" json:"number"`

	Book *Book `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE"`
}
