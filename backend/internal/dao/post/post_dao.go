package postdao

import (
	"hackathon/db/sqlc/generated"
)

type PostDAO struct {
	db *sqlc.Queries
}

func NewPostDAO(db *sqlc.Queries) *PostDAO {
	return &PostDAO{db: db}
}

