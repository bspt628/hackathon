package likedao

import (
	"hackathon/db/sqlc/generated"
	"database/sql"
)

type LikeDAO struct {
	db      *sql.DB       // データベース接続
	queries *sqlc.Queries // クエリ生成
}

func NewLikeDAO(db *sql.DB) *LikeDAO {
	return &LikeDAO{
		db:      db,
		queries: sqlc.New(db),
	}
}

