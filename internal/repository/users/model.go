package users

import (
	"github.com/zigal0/sport_bot/internal/domain"
)

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
}

func (u *User) ToDomain() domain.User {
	return domain.User{
		ID:       u.ID,
		Username: u.Username,
	}
}

func UserSliceToDomain(dbUsers []User) []domain.User {
	domainUsers := make([]domain.User, 0, len(dbUsers))

	for _, dbUser := range dbUsers {
		domainUsers = append(domainUsers, dbUser.ToDomain())
	}

	return domainUsers
}

func DBUserFromDomain(domainUser domain.User) User {
	return User{
		ID:       domainUser.ID,
		Username: domainUser.Username,
	}
}

func DBUserSliceFromDomain(domainUsers []domain.User) []User {
	dbUsers := make([]User, 0, len(domainUsers))

	for _, domainUser := range domainUsers {
		dbUsers = append(dbUsers, DBUserFromDomain(domainUser))
	}

	return dbUsers
}
