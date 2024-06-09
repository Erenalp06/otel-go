package database

import (
	"fmt"
	"os"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	
	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("missing one or more environment variables: DB_HOST=%s, DB_PORT=%s, DB_USER=%s, DB_PASSWORD=%s, DB_NAME=%s", host, port, user, password, dbname)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=Europe/Istanbul",
		host,
		port,
		user,
		password,
		dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithAttributes(attribute.String("DB_HOST", os.Getenv("DB_HOST"))))); err != nil {
		panic(err)
	}
	return db, nil
}
