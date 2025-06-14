package config

import (
	"API_GO/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	//Check if DB file exists

	dbPath := "./db/main.db"
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("dabase file not found, creating...")
		//crear el archivo de base de datos y directorio
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
	//CREAR DB y coneccion
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.ErrorF("sqlite opening ERROR:", err)
		return nil, err
	}
	//migrate the schema
	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		logger.ErrorF("sqlite automigration ERROR:", err)
		return nil, err
	}

	//return DB
	return db, nil
}
