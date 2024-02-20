package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// depois configurar o BD do outro jeito, pois fica melhor para injeção de dependencia
var (
	//DB is an instance of the data base
	DB *gorm.DB
)

// ConectDB starts the data base
func ConectDB() (*gorm.DB, error) {
	strConection := "{usuario}:{senha}@/{seu-bd}?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(strConection), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CloseDB close db conection
func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
