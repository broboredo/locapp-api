package handler

import (
	"github.com/broboredo/locapp-api/config"
	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	Db     *gorm.DB
)

func InitHandler() {
	Logger = config.GetLogger("handler")
	Db = config.GetDB()
}
