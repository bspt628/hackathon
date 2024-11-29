package followdao

import (
	"hackathon/db/sqlc/generated"
)

type FollowDAO struct {
	db *sqlc.Queries
}

func NewFollowDAO(db *sqlc.Queries) *FollowDAO {
	return &FollowDAO{db: db}
}

