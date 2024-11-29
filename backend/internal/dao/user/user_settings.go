package userdao

import (
	"context"
	"hackathon/db/sqlc/generated"
)

// プロフィール情報を更新
func (dao *UserDAO) UpdateUserSettings(ctx context.Context, params sqlc.UpdateUserSettingsParams) error {
	// SQLC を使用してデータベースにアクセス
	return dao.queries.UpdateUserSettings(ctx, params)
}
