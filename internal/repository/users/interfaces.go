package users

import (
	"github.com/zigal0/sport_bot/internal/model"
)

type UsersRepo interface {
	AddUser(user model.User) error
}
