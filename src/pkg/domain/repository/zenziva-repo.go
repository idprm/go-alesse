package repository

import "github.com/idprm/go-alesse/src/pkg/domain/model"

type ZenzivaRepository interface {
	GetByID(id int) (model.Zenziva, error)
	GetAll() ([]model.Zenziva, error)
	Save(model.Zenziva) error
}
