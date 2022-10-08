package database

import (
	"fmt"

	"github.com/idprm/go-alesse/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg config.DBCfg) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.GetDBUser(),
		cfg.GetDBPass(),
		cfg.GetDBHost(),
		cfg.GetDBPort(),
		cfg.GetDBName())),
		&gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, err
}
