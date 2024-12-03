package postdao

import (
	"hackathon/db/sqlc/generated"
	"database/sql"
)

type PostDAO struct {
	db      *sql.DB       // データベース接続
	queries *sqlc.Queries // クエリ生成
}

func NewPostDAO(db *sql.DB) *PostDAO {
	return &PostDAO{
		db:      db,
		queries: sqlc.New(db),
	}
}

