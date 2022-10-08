package service

import (
	"github.com/idprm/go-alesse/src/domain/model"
	"github.com/idprm/go-alesse/src/domain/repository"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Conn *gorm.DB
}

func NewUserRepositoryWithRDB(conn *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{Conn: conn}
}

func (r *UserRepositoryImpl) Get(id int) (*model.User, error) {
	user := &model.User{}
	if err := r.Conn.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetAll() ([]model.User, error) {
	users := []model.User{}
	if err := r.Conn.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetBySlug(slug string) (*model.User, error) {
	user := &model.User{}
	if err := r.Conn.Where("username", slug).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) Save(user *model.User) error {
	if err := r.Conn.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) Remove(id int) error {
	tx := r.Conn.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	user := model.User{}
	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *UserRepositoryImpl) Update(user *model.User) error {
	if err := r.Conn.Model(&user).UpdateColumns(model.User{
		Name:    user.Name,
		Dob:     user.Dob,
		Gender:  user.Gender,
		Address: user.Address,
	}).Error; err != nil {
		return nil
	}
	return nil
}
