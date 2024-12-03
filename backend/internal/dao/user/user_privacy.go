package userdao

import (
	"context"
	sqlc "hackathon/db/sqlc/generated"
)

// プライバシー設定を更新
func (dao *UserDAO) UpdateUserPrivacy(ctx context.Context, params sqlc.UpdateUserPrivacyParams) error {
	// SQLC を使用してデータベースにアクセス
	return dao.queries.UpdateUserPrivacy(ctx, params)
}
