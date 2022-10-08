package database

import (
	"log"

	"github.com/idprm/go-alesse/src/config"
	"github.com/idprm/go-alesse/src/domain/model"
	"gorm.io/gorm"
)

func DBMigrate() (*gorm.DB, error) {
	var cfg config.DBCfg
	conn, err := NewDB(cfg)
	if err != nil {
		return nil, err
	}

	conn.AutoMigrate(model.Doctor{})
	log.Println("Migration has been processed")

	return conn, nil

}
