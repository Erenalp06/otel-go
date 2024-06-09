package migrations

import (
	"github.com/Erenalp06/otel-go/pkg/models"
	"gorm.io/gorm"
)

func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	return err
}
