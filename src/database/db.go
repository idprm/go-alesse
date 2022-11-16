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
		&model.SuperAdmin{},
		&model.Admin{},
		&model.Disease{},
		&model.Medicine{},
		&model.User{},
		&model.Verify{},
		&model.Doctor{},
		&model.Specialist{},
		&model.Officer{},
		&model.Apothecary{},
		&model.Courier{},
		&model.Driver{},
		&model.Chat{},
		&model.MedicalResume{},
		&model.Referral{},
		&model.Prescription{},
		&model.PrescriptionMedicine{},
		&model.Homecare{},
		&model.HomecareMedicine{},
		&model.HomecareOfficer{},
		&model.HomecareUpload{},
		&model.Pharmacy{},
		&model.Treatment{},
		&model.Shipment{},
		&model.Feedback{},
		&model.Sendbird{},
		&model.Zenziva{},
	)

	// TODO: Add seeders
	var config []model.Config
	var healthcenter []model.Healthcenter
	var disease []model.Disease
	var medicine []model.Medicine
	var doctor []model.Doctor
	var specialist []model.Specialist
	var officer []model.Officer
	var driver []model.Driver
	var apothecary []model.Apothecary
	var courier []model.Courier
	var admin []model.Admin
	var superadmin []model.SuperAdmin

	resultConfig := db.Find(&config)
	resultHealthCenter := db.Find(&healthcenter)
	resultDisease := db.Find(&disease)
	resultMedicine := db.Find(&medicine)
	resultDoctor := db.Find(&doctor)
	resultSpecialist := db.Find(&specialist)
	resultOfficer := db.Find(&officer)
	resultDriver := db.Find(&driver)
	resultApothecary := db.Find(&apothecary)
	resultCourier := db.Find(&courier)
	resultAdmin := db.Find(&admin)
	resultSuperadmin := db.Find(&superadmin)

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

	if resultSpecialist.RowsAffected == 0 {
		for i, _ := range specialists {
			db.Model(&model.Specialist{}).Create(&specialists[i])
		}
	}

	if resultOfficer.RowsAffected == 0 {
		for i, _ := range officers {
			db.Model(&model.Officer{}).Create(&officers[i])
		}
	}

	if resultDriver.RowsAffected == 0 {
		for i, _ := range drivers {
			db.Model(&model.Driver{}).Create(&drivers[i])
		}
	}

	if resultApothecary.RowsAffected == 0 {
		for i, _ := range apothecaries {
			db.Model(&model.Apothecary{}).Create(&apothecaries[i])
		}
	}

	if resultCourier.RowsAffected == 0 {
		for i, _ := range couriers {
			db.Model(&model.Courier{}).Create(&couriers[i])
		}
	}

	if resultAdmin.RowsAffected == 0 {
		for i, _ := range admins {
			db.Model(&model.Admin{}).Create(&admins[i])
		}
	}

	if resultSuperadmin.RowsAffected == 0 {
		for i, _ := range superadmins {
			db.Model(&model.SuperAdmin{}).Create(&superadmins[i])
		}
	}

	Datasource = &NewDatasource{db: db, sqlDb: sqlDb}
}
