package followdao

import (
	"hackathon/db/sqlc/generated"
	"database/sql"
)

type FollowDAO struct {
	db      *sql.DB       // データベース接続
	queries *sqlc.Queries // クエリ生成
}

func NewFollowDAO(db *sql.DB) *FollowDAO {
	return &FollowDAO{
		db:      db,
		queries: sqlc.New(db),
	}
}
