package repository

import "github.com/idprm/go-alesse/src/pkg/domain/model"

type MedicalRepository interface {
	Get(id int) (model.Medical, error)
	GetAll() ([]model.Medical, error)
	Add(model.Medical) error
	Remove(id int) error
	Update(model.Medical) error
}
