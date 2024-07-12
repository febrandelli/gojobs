package config

import (
	"github.com/febrandelli/gojobs/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("Creating database at %s", dbPath)
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error opening SQLite database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schema.Opening{})
	if err != nil {
		logger.Errorf("Error auto-migrating SQLite database: %v", err)
		return nil, err
	}

	return db, err
}
