package repository

import "github.com/idprm/go-alesse/src/pkg/domain/model"

type UserRepository interface {
	GetByID(id int) (model.User, error)
	GetAll() ([]model.User, error)
	GetBySlug(slug string) (model.User, error)
	Save(model.User) error
	Remove(id int) error
	Update(model.User) error
}