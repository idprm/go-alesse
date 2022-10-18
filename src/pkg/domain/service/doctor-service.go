package service

import "github.com/idprm/go-alesse/src/pkg/domain/model"

type DoctorService interface {
	Add(model.Doctor) error
	GetAll() ([]model.Doctor, error)
}
