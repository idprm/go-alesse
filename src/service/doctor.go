package service

import "github.com/idprm/go-alesse/src/domain/repository"

type doctorService struct {
	doctorRepo repository.DoctorRepository
}

func NewDoctorService(doctorRepo repository.DoctorRepository) *doctorService {
	return &doctorService{
		doctorRepo: doctorRepo,
	}
}
