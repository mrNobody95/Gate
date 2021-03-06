package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"syscall"
)

var db *gorm.DB
var err error

type DatabaseConfig struct {
	Type     DbType
	Username string
	Password string
	CharSet  string
	Name     string
	Host     string
	Port     int
	SSLMode  bool
	Timezone syscall.Timezoneinformation
}

type DbType string

const (
	MYSQL      DbType = "mysql"
	SQLITE     DbType = "sqlite"
	PostgreSQL DbType = "postgre"
)

func (conf *DatabaseConfig) Initialize(models ...interface{}) error {
	switch conf.Type {
	case MYSQL:
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Name)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
	case PostgreSQL:
		ssl := "disable"
		if conf.SSLMode {
			ssl = "enable"
		}
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=UTC", conf.Host, conf.Username, conf.Password, conf.Name, conf.Port, ssl)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return err
		}
	case SQLITE:
		db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	}

	return migrate(models)
}

func migrate(models ...interface{}) error {
	for _, model := range models {
		if err := db.AutoMigrate(&model); err != nil {
			return err
		}
	}
	//if err := db.AutoMigrate(&Candle{}); err != nil {
	//	return err
	//}
	//if err := db.AutoMigrate(&Order{}); err != nil {
	//	return err
	//}

	return nil
}

func GetDB() *gorm.DB {
	return db
}
