package command_processor

import (
	"github.com/zigal0/sport_bot/internal/domain"
)

type UserRepo interface {
	GetUser(id int64) (domain.User, error)
	AddUser(user domain.User) error
	UpdateUser(user domain.User) error
	DeleteUser(id int64) error
}
