package repository

import "github.com/idprm/go-alesse/src/pkg/domain/model"

type VerifyRepository interface {
	GetByID(id int) (model.Verify, error)
	GetAll() ([]model.Verify, error)
	Add(model.Verify) error
	Remove(id int) error
}
