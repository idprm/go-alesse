package repository

import "github.com/idprm/go-alesse/src/domain/model"

type SendbirdRepository interface {
	Get(id int) (*model.Sendbird, error)
	GetAll() ([]model.Sendbird, error)
	Save(*model.Sendbird) error
}
