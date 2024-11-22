package usecase

import (
	"context"
	"database/sql"
	sqlc "hackathon/db/sqlc/generated"
	"hackathon/domain"
)

func (uc *UserUsecase) UpdateUserBanStatus(ctx context.Context, isBanned bool, id string) (*domain.UserBanStatusUpdateResult, error) {
	// UpdateUserProfileParams構造体にデータをセット
	arg := sqlc.UpdateUserBanStatusParams{
		IsBanned: 		sql.NullBool{Bool: isBanned, Valid: true},
		ID: 			id,
	}

	// ユーザー情報を更新
	err := uc.dao.UpdateUserBanStatus(ctx, arg)
	if err != nil {
		return nil, err
	}

	return domain.NewUserBanStatusUpdateResult(isBanned), nil
}
