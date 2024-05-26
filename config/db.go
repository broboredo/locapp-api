package config

import (
	"fmt"
	"github.com/broboredo/locapp-api/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	logger = GetLogger("db")

	dbUser := "root"
	dbPass := "root"
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "locapp"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	logger.Info("starting db...")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errorf("db starting error: %v", err)
		return nil, err
	}
	logger.Info("db started")

	logger.Info("init db migrating...")
	err = db.AutoMigrate(&schemas.Product{})
	err = db.AutoMigrate(&schemas.Customer{})

	if err != nil {
		logger.Errorf("db migrating error: %v", err)
		return nil, err
	}
	logger.Info("db migrated successfully")

	return db, nil
}
