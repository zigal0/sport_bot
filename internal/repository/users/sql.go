package users

const (
	QueryGetUser = `
SELECT id, username FROM users WHERE id = $1
`

	QueryInsertUser = `
INSERT INTO users (id, username) VALUES (:id, :username)
`

	QueryUpdateUser = `
UPDATE users SET username = :username WHERE id = :id
`

	QueryDeleteUser = `
DELETE FROM users WHERE id = :id
`
)
