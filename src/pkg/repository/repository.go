package repository

import "github.com/idprm/go-alesse/src/pkg/model"

type NotifRepository interface {
	GetAll() ([]model.Notif, error)
	GetByID(id uint64) (model.Notif, error)
	Add(notif model.Notif) error
	// Update(notif model.Notif) error
	Delete(id uint64) error
}
