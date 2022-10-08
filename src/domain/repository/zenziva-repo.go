package repository

import "github.com/idprm/go-alesse/src/domain/model"

type ZenzivaRepository interface {
	Get(id int) (*model.Zenziva, error)
	GetAll() ([]model.Zenziva, error)
	Save(*model.Zenziva) error
}
