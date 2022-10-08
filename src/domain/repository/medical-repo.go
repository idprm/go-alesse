package repository

import "github.com/idprm/go-alesse/src/domain/model"

type MedicalRepository interface {
	Get(id int) (*model.Medical, error)
	GetAll() ([]model.Medical, error)
	Save(*model.Medical) error
	Remove(id int) error
	Update(*model.Medical) error
}
