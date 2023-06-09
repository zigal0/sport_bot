package users

import (
	"github.com/zigal0/sport_bot/internal/domain"
)

type UsersRepo interface {
	AddUser(user domain.User) error
}
