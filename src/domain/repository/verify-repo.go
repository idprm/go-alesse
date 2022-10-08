package repository

import "github.com/idprm/go-alesse/src/domain/model"

type VerifyRepository interface {
	Get(id int) (*model.Verify, error)
	GetAll() ([]model.Verify, error)
	Save(*model.Verify) error
	Remove(id int) error
}
