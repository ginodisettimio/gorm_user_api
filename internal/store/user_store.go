package store

import (
	"GOrm/internal/model"

	"gorm.io/gorm"
)

type Store interface {
	GetAll() ([]*model.User, error)
	GetById(id int) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(id int, user *model.User) (*model.User, error)
	Delete(id int) error
}

type UserStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) Store {
	return &UserStore{db: db}
}

func (u *UserStore) GetAll() ([]*model.User, error) {
	var users []*model.User
	result := u.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u *UserStore) GetById(id int) (*model.User, error) {
	var user *model.User
	result := u.db.Find(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *UserStore) Create(user *model.User) (*model.User, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *UserStore) Update(id int, user *model.User) (*model.User, error) {
	result := u.db.First(&user, id).Updates(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *UserStore) Delete(id int) error {
	result := u.db.Delete(model.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
