package persistence

import (
	"github.com/idprm/go-alesse/src/pkg/model"
	"gorm.io/gorm"
)

type NotifRepo struct {
	db *gorm.DB
}

func NewNotifRepository(db *gorm.DB) *NotifRepo {
	return &NotifRepo{db}
}

func (r *NotifRepo) GetAll() ([]model.Notif, error) {
	var notifs []model.Notif
	err := r.db.Debug().Find(&notifs).Error
	if err != nil {
		return nil, err
	}
	return notifs, nil
}

func (r *NotifRepo) GetByID(id uint64) (model.Notif, error) {
	var notif model.Notif
	err := r.db.Debug().Where("id = ?", id).Take(&notif).Error
	if err != nil {
		return notif, err
	}
	return notif, nil
}

func (r *NotifRepo) Add(notif model.Notif) error {
	err := r.db.Debug().Create(&notif).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *NotifRepo) Delete(id uint64) error {
	var notif model.Notif
	err := r.db.Debug().Where("id = ?", id).Delete(&notif).Error
	if err != nil {
		return err
	}
	return nil
}
