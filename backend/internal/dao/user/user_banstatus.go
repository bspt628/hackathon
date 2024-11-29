package userdao

import (
	"context"
	"hackathon/db/sqlc/generated"
)

// プライバシー設定を更新
func (dao *UserDAO) UpdateUserBanStatus(ctx context.Context, params sqlc.UpdateUserBanStatusParams) error {
	return dao.db.UpdateUserBanStatus(ctx, params)
}
