package migrations

import (
	"github.com/boazsoftwares/Boaz.Api.Go/models"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(models.Empresa)
}
