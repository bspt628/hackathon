package dao

import (
	"context"
	"hackathon/db/sqlc/generated"
)

func (dao *UserDAO) UpdateUserNotifications(ctx context.Context, params sqlc.UpdateUserNotificationsParams) error {
	// SQLC を使用してデータベースにアクセス
	return dao.db.UpdateUserNotifications(ctx, params)
}