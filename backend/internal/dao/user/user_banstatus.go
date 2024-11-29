package dao

import (
	"context"
	"hackathon/db/sqlc/generated"
)

// プライバシー設定を更新
func (dao *UserDAO) UpdateUserBanStatus(ctx context.Context, params sqlc.UpdateUserBanStatusParams) error {
	// SQLC を使用してデータベースにアクセス
	return dao.db.UpdateUserBanStatus(ctx, params)
}