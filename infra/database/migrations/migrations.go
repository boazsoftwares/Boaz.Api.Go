package migrations

import (
	entities "github.com/boazsoftwares/Boaz.Api.Go/domain/entities"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(entities.Company{})
}
