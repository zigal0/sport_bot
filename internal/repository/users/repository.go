package users

import (
	"github.com/zigal0/sport_bot/internal/model"
)

type usersRepo struct{}

func NewUsersRepo() UsersRepo {
	return &usersRepo{}
}

func (r *usersRepo) AddUser(user model.User) error {
	return nil
}
