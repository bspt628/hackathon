package usecase

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"hackathon/domain"
)

func (uc *UserUsecase) UpdateUserPrivacy(ctx context.Context, isPrivate bool, id string) (*domain.UserPrivacyUpdateResult, error) {


	// 更新結果をまとめる

	// UpdateUserProfileParams構造体にデータをセット
	arg := sqlc.UpdateUserPrivacyParams{
		IsPrivate: 		sql.NullBool{Bool: isPrivate, Valid: true},
		ID: 			id,
	}

	// ユーザー情報を更新
	err := uc.dao.UpdateUserPrivacy(ctx, arg)
	if err != nil {
		return nil, err
	}

	return domain.NewUserPrivacyUpdateResult(isPrivate), nil
}
