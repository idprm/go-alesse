package persistence

import (
	"fmt"

	"github.com/idprm/go-alesse/src/pkg/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repositories struct {
	Notif repository.NotifRepository
	db    *gorm.DB
}

func NewRepositories(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*Repositories, error) {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Notif: NewNotifRepository(db),
		db:    db,
	}, nil
}
