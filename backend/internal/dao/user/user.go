package userdao

import sqlc "hackathon/db/sqlc/generated"

type UserDAO struct {
	db *sqlc.Queries
}

func NewUserDAO(db *sqlc.Queries) *UserDAO {
	return &UserDAO{db: db}
}
