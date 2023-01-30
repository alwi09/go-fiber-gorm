package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	const MYSQL = "alwi09:alwiirfani11@tcp(127.0.0.1:3306)/rest_full_api_todolist?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	return db, nil
}
