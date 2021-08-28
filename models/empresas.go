package models

import (
	"time"

	"gorm.io/gorm"
)

type Empresa struct {
	ID          uint           `json:"Id" gorm:"column:Id;not null;primaryKey;autoIncrement"`
	Cnpj        string         `json:"Cnpj" gorm:"column:Cnpj;size:14;index;not null;unique"`
	RazaoSocial string         `json:"RazaoSocial" gorm:"column:RazaoSocial;size:100;z;not null;unique"`
	NomeFatasia string         `json:"NomeFantasia" gorm:"column:NomeFantasia;size:100;not null;unique"`
	CreatedAt   time.Time      `json:"CreatedAt" gorm:"column:CreatedAt;not null;type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `json:"UpdatedAt" gorm:"column:UpdatedAt;type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"column:DeletedAt"`
}

// Set User's table name to be `profiles`
func (Empresa) TableName() string {
	return "Empresa"
}

// type Empresa struct {
// 	ID          uint           `json:"Id" gorm:"column:Id;not null;primaryKey;autoIncrement"`
// 	Cnpj        string         `json:"Cnpj" gorm:"column:Cnpj;size:14;uniqueIndex;index;not null"`
// 	RazaoSocial string         `json:"RazaoSocial" gorm:"column:RazaoSocial;size:100;uniqueIndex;not null"`
// 	NomeFatasia string         `json:"NomeFantasia" gorm:"column:NomeFantasia;size:100;not null"`
// 	CreatedAt   time.Time      `json:"CreatedAt" gorm:"column:CreatedAt;not null"`
// 	UpdatedAt   time.Time      `json:"UpdatedAt" gorm:"column:UpdatedAt;"`
// 	DeletedAt   gorm.DeletedAt `gorm:"index;column:DeletedAt;"`
// }
