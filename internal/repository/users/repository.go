package users

import (
	"github.com/jmoiron/sqlx"
	"github.com/zigal0/sport_bot/internal/domain"
	"github.com/pkg/errors"
)

type usersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) UsersRepo {
	return &usersRepo{
		db: db,
	}
}

func (r *usersRepo) AddUser(user domain.User) error {
	
	dbUser := DBUserFromDomain(user)

	_, err := r.db.NamedExec(QueryInsertUser, dbUser)
	if err != nil {
		return errors.Wrap(err, "usersRepo.AddUser")
	}

	return nil
}
