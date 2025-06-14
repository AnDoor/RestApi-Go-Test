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

	//Inicializar SQLITE
	db, err = InitializeSQLite()

	if err != nil {
		return fmt.Errorf("ERROR: error initialisin sqlite:%v", err)
	}
	return nil
}

func GetSQLITE() *gorm.DB {
	return db
}
func GetLogger(p string) *Logger {
	//inici logger
	logger = NewLogger(p)
	return logger
}
