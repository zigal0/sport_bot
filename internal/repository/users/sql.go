package users

const (
	QueryInsertUser = `
INSERT INTO users (id, username) VALUES (:id, :username)
`
)