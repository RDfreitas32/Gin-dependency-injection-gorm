package db

import (
	models "dependency/pkg/moldels"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// depois configurar o BD do outro jeito, pois fica melhor para injeção de dependencia
var (
	//DB is an instance of the data base
	DB *gorm.DB
)

// InitDB starts the data base
func InitDB() error {
	var err error
	DB, err = gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto migration
	err = DB.AutoMigrate(&models.Task{})
	if err != nil {
		return err
	}
	return nil
}
