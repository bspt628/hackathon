package dao

import (
	"context"
	"hackathon/db/sqlc/generated"
)

// プライバシー設定を更新
func (dao *UserDAO) UpdateUserPrivacy(ctx context.Context, params sqlc.UpdateUserPrivacyParams) error {
	// SQLC を使用してデータベースにアクセス
	return dao.db.UpdateUserPrivacy(ctx, params)
}