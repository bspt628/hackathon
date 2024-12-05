package repostdao

import (
	"hackathon/db/sqlc/generated"
	"database/sql"
)

type RepostDAO struct {
	db      *sql.DB       // データベース接続
	queries *sqlc.Queries // クエリ生成
}

func NewRepostDAO(db *sql.DB) *RepostDAO {
	return &RepostDAO{
		db:      db,
		queries: sqlc.New(db),
	}
}
