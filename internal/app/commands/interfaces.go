package commands

import (
	"github.com/zigal0/sport_bot/internal/domain"
)

type CommandProcessor interface{
	GetUser(id int64) (domain.User, error)
	AddUser(user domain.User) error
	UpdateUser(user domain.User) error
	DeleteUser(id int64) error
}
