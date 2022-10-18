package sql

import (
	"context"

	"github.com/idprm/go-alesse/src/pkg/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	r := &UserRepository{
		DB: db,
	}
	return r
}

func (r UserRepository) Save(ctx context.Context, user *model.User) error {
	if err := r.DB.Save(user).Find(&user).Error; err != nil {
		return err
	}
	return nil
}
