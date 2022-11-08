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
		&model.Disease{},
		&model.Medicine{},
		&model.User{},
		&model.Verify{},
		&model.Doctor{},
		&model.Officer{},
		&model.Chat{},
		&model.MedicalResume{},
		&model.Referral{},
		&model.Prescription{},
		&model.PrescriptionMedicine{},
		&model.Homecare{},
		&model.HomecareMedicine{},
		&model.Sendbird{},
		&model.Zenziva{},
	)

	// TODO: Add seeders
	var config []model.Config
	var healthcenter []model.Healthcenter
	var disease []model.Disease
	var medicine []model.Medicine
	var doctor []model.Doctor
	var officer []model.Officer

	resultConfig := db.Find(&config)
	resultHealthCenter := db.Find(&healthcenter)
	resultDisease := db.Find(&disease)
	resultMedicine := db.Find(&medicine)
	resultDoctor := db.Find(&doctor)
	resultOfficer := db.Find(&officer)

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

	if resultDisease.RowsAffected == 0 {
		for i, _ := range diseases {
			db.Model(&model.Disease{}).Create(&diseases[i])
		}
	}

	if resultMedicine.RowsAffected == 0 {
		for i, _ := range medicines {
			db.Model(&model.Medicine{}).Create(&medicines[i])
		}
	}

	if resultDoctor.RowsAffected == 0 {
		for i, _ := range doctors {
			db.Model(&model.Doctor{}).Create(&doctors[i])
		}
	}

	if resultOfficer.RowsAffected == 0 {
		for i, _ := range officers {
			db.Model(&model.Officer{}).Create(&officers[i])
		}
	}

	Datasource = &NewDatasource{db: db, sqlDb: sqlDb}
}
