package models

import (
	"time"

	"gorm.io/gorm"
)

type Empresa struct {
	Id          uint           `json:"Id" gorm:"primaryKey"`
	Cnpj        string         `json:"Cnpj"`
	RazaoSocial string         `json:"RazaoSocial"`
	NomeFatasia string         `json:"NomeFantasia"`
	CreatedAt   time.Time      `json:"Created"`
	UpdatedAt   time.Time      `json: "Updated"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"Deleted"`
}
