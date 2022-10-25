package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/idprm/go-alesse/src/pkg/model"
	"github.com/idprm/go-alesse/src/pkg/util/localconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Datasource *NewDatasource

type NewDatasource struct {
	db    *gorm.DB
	sqlDb *sql.DB
}

func (d NewDatasource) DB() *gorm.DB {
	return d.db
}

func (d NewDatasource) SqlDB() *sql.DB {
	return d.sqlDb
}

func Connect() {

	var db *gorm.DB
	var sqlDb *sql.DB

	secret, err := localconfig.LoadSecret("src/server/secret.yaml")
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		secret.DB.UserName,
		secret.DB.Password,
		secret.DB.Host,
		secret.DB.Port,
		secret.DB.DBName,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err.Error())
		panic("Could not connect with the database!")
	}

	sqlDb, _ = db.DB()
	sqlDb.SetConnMaxLifetime(time.Minute * 2)
	sqlDb.SetMaxOpenConns(10000)
	sqlDb.SetMaxIdleConns(10000)

	// try to establish connection
	if sqlDb != nil {
		err = sqlDb.Ping()
		if err != nil {
			log.Fatal("cannot connect to db:", err.Error())
		}
	}

	log.Println("Connected to database successfully")

	db.AutoMigrate(
		&model.Config{},
		&model.Healthcenter{},
		&model.User{},
		&model.Verify{},
		&model.Doctor{},
		&model.Clinic{},
		&model.Chat{},
		&model.Medical{},
		&model.Sendbird{},
		&model.Zenziva{},
	)

	// TODO: Add seeders
	var config []model.Config
	var healthcenter []model.Healthcenter
	var doctor []model.Doctor
	var clinic []model.Clinic

	resultConfig := db.Find(&config)
	resultHealthCenter := db.Find(&healthcenter)
	resultDoctor := db.Find(&doctor)
	resultClinic := db.Find(&clinic)

	if resultConfig.RowsAffected == 0 {
		for i, _ := range configs {
			db.Model(&model.Config{}).Create(&configs[i])
		}
	}

	if resultHealthCenter.RowsAffected == 0 {
		for i, _ := range healthcenters {
			db.Model(&model.Healthcenter{}).Create(&healthcenters[i])
		}
	}

	if resultDoctor.RowsAffected == 0 {
		for i, _ := range doctors {
			db.Model(&model.Doctor{}).Create(&doctors[i])
		}
	}

	if resultClinic.RowsAffected == 0 {
		for i, _ := range clinics {
			db.Model(&model.Clinic{}).Create(&clinics[i])
		}
	}

	Datasource = &NewDatasource{db: db, sqlDb: sqlDb}
}
