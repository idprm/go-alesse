package repository

import "github.com/idprm/go-alesse/src/domain/model"

type UserRepository interface {
	Get(id int) (*model.User, error)
	GetAll() ([]model.User, error)
	GetBySlug(slug string) (*model.User, error)
	Save(*model.User) error
	Remove(id int) error
	Update(*model.User) error
}
