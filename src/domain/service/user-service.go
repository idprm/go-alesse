package service

import "github.com/idprm/go-alesse/src/domain/model"

type UserService interface {
	Add(model.User) error
	GetByID(userID int) (model.User, error)
	GetAll() ([]model.User, error)
}
