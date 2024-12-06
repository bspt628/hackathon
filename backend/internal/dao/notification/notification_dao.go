package notificationdao

import (
	"hackathon/db/sqlc/generated"
	"database/sql"
)

type NotificationDAO struct {
	db      *sql.DB       // データベース接続
	queries *sqlc.Queries // クエリ生成
}

func NewNotificationDAO(db *sql.DB) *NotificationDAO {
	return &NotificationDAO{
		db:      db,
		queries: sqlc.New(db),
	}
}

