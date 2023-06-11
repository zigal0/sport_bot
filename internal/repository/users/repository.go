package users

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/zigal0/sport_bot/internal/domain"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{
		db: db,
	}
}

func (r *UsersRepo) GetUser(id int64) (domain.User, error) {
	user := User{}

	err := r.db.Get(
		&user,
		QueryGetUser,
		id,
	)
	if err != nil {
		return domain.User{}, errors.Wrap(err, "UsersRepo.GetUser")
	}

	return user.ToDomain(), nil
}

func (r *UsersRepo) AddUser(user domain.User) error {
	dbUser := DBUserFromDomain(user)

	_, err := r.db.NamedExec(QueryInsertUser, dbUser)
	if err != nil {
		return errors.Wrap(err, "UsersRepo.AddUser")
	}

	return nil
}

func (r *UsersRepo) UpdateUser(user domain.User) error {
	dbUser := DBUserFromDomain(user)

	_, err := r.db.NamedExec(QueryUpdateUser, dbUser)
	if err != nil {
		return errors.Wrap(err, "UsersRepo.UpdateUser")
	}

	return nil
}

func (r *UsersRepo) DeleteUser(id int64) error {
	_, err := r.db.NamedExec(
		QueryDeleteUser,
		map[string]interface{}{"id": id},
	)
	if err != nil {
		return errors.Wrap(err, "UsersRepo.DeleteUser")
	}

	return nil
}
