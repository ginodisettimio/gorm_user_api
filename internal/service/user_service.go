package service

import (
	"GOrm/internal/model"
	"GOrm/internal/store"
	"errors"
	"regexp"
)

type UserService struct {
	store store.Store
}

func New(store store.Store) *UserService {
	// TODO: Implementar logger
	return &UserService{store: store}
}

func (u *UserService) GetAllUsers() ([]*model.User, error) {
	// TODO: Implementar logger
	return u.store.GetAll()
}

func (u *UserService) GetUserById(id int) (*model.User, error) {
	// TODO: Implementar logger
	return u.store.GetById(id)
}

func (u *UserService) CreateUser(user *model.User) (*model.User, error) {
	// TODO: Implementar logger
	if user.Name == "" {
		return nil, errors.New("Nombre inválido")
	}

	if !isValidPassword(user.Password) {
		return nil, errors.New("Contraseña inválida")
	}

	return u.store.Create(user)
}

func (u *UserService) UpdateUser(id int, user *model.User) (*model.User, error) {
	// TODO: Implementar logger
	oldUser, err := u.store.GetById(id)
	if err != nil {
		return nil, err
	}

	if user.Name == "" {
		user.Name = oldUser.Name
	}
	if !isValidPassword(user.Password) {
		user.Password = oldUser.Password
	}

	return u.store.Update(id, user)
}

func (u *UserService) DeleteUser(id int) error {
	// TODO: Implementar logger
	return u.store.Delete(id)
}

func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`)
	hasSpecial := regexp.MustCompile(`[^A-Za-z0-9\s]`)

	if !hasUpper.MatchString(password) && !hasSpecial.MatchString(password) {
		return false
	}

	return true
}
