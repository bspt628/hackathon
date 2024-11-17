package dao

import (
	"hackathon/db/sqlc/generated"
)

type UserDAO struct {
	db *db.Queries
}

func NewUserDAO(db *db.Queries) *UserDAO {
	return &UserDAO{db: db}
}

