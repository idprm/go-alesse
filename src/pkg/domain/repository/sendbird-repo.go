package repository

import "github.com/idprm/go-alesse/src/pkg/domain/model"

type SendbirdRepository interface {
	GetByID(id int) (model.Sendbird, error)
	GetAll() ([]model.Sendbird, error)
	Add(model.Sendbird) error
}
