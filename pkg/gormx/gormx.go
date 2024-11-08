package gormx

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	Type string `ini:"type"`
	DSN  string `ini:"dsn"`
}

func New(cfg Config) (db *gorm.DB, err error) {
	switch cfg.Type {
	case "sqlite3":
		// dsn := "exe.db"
		db, err = gorm.Open(sqlite.Open(cfg.DSN), &gorm.Config{})
	case "mysql":
		// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{})
	default:
		err = errors.New("mode not supported")
	}
	return
}
