package userdao

import (
	"hackathon/db/sqlc/generated"
	"database/sql"
)

type UserDAO struct {
	db      *sql.DB       // データベース接続
	queries *sqlc.Queries // クエリ生成
}

// コンストラクタ
func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{
		db:      db,
		queries: sqlc.New(db),
	}
}
