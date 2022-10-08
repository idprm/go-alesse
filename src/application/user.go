package application

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/idprm/go-alesse/src/config"
	"github.com/idprm/go-alesse/src/database"
	"github.com/idprm/go-alesse/src/domain/model"

	"github.com/idprm/go-alesse/src/service"
)

func GetUser(id int) (*model.User, error) {
	var cfg config.DBCfg
	conn, err := database.NewDB(cfg)

	if err != nil {
		return nil, err
	}

	repo := service.NewUserRepositoryWithRDB(conn)
	return repo.Get(id)
}

func GetAllUsers(limit int, page int) ([]model.User, error) {
	var cfg config.DBCfg
	_, err := database.NewDB(cfg)

	if err != nil {
		return nil, err
	}

	var users []model.User
	pagination.Paging(&pagination.Param{
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &users)

	return users, nil
}

func AddUser(p model.User) error {
	var cfg config.DBCfg
	conn, err := database.NewDB(cfg)
	if err != nil {
		return err
	}

	repo := service.NewUserRepositoryWithRDB(conn)
	return repo.Save(&p)
}

func RemoveUser(id int) error {
	var cfg config.DBCfg
	conn, err := database.NewDB(cfg)
	if err != nil {
		return err
	}

	repo := service.NewUserRepositoryWithRDB(conn)
	return repo.Remove(id)
}

func UpdateUser(p model.User, id int) error {
	var cfg config.DBCfg
	conn, err := database.NewDB(cfg)
	if err != nil {
		return err
	}

	repo := service.NewUserRepositoryWithRDB(conn)
	p.ID = uint64(id)
	return repo.Update(&p)
}
