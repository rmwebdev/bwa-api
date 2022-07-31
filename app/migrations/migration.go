package migrations

import (
	"api/app/model"

	"gorm.io/gorm"
)

// AutoMigrate all tables
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(ModelMigrations...)
}

// ModelMigrations models to migrate
var ModelMigrations []interface{} = []interface{}{
	&model.User{},
}
