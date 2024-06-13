package config

import (
	"fmt"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	logger := NewLogger("config")

	logger.Info("initializing db...")
	db, err = InitDB()

	if err != nil {
		return fmt.Errorf("error initializing db: %v", err)
	}

	return nil
}
func GetDB() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger := NewLogger(prefix)
	return logger
}

func LoadEnv() {
	loadEnv()
}
