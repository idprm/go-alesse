package repository

import "github.com/idprm/go-alesse/src/domain/model"

type DoctorRepository interface {
	Get(id int) (*model.Doctor, error)
	GetAll() ([]model.Doctor, error)
	Save(*model.Doctor) error
	Remove(id int) error
	Update(*model.Doctor) error
}
