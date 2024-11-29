package userdao

import (
	"context"
	"hackathon/db/sqlc/generated"
)

// プロフィール情報を更新
func (dao *UserDAO) UpdateUserProfile(ctx context.Context, params sqlc.UpdateUserProfileParams) error {
	// SQLC を使用してデータベースにアクセス
	return dao.queries.UpdateUserProfile(ctx, params)
}
