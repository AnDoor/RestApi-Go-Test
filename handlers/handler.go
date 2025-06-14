package handlers

import (
	"API_GO/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitHandler(){
	logger = config.GetLogger("handler")
	db = config.GetSQLITE()
	
}