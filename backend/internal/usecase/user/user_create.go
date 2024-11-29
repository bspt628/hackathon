package userusecase

import (
	"context"
	"database/sql"
	sqlc "hackathon/db/sqlc/generated"
)

func (uc *UserUsecase) CreateUser(ctx context.Context, email, password, username, displayName string) (*sqlc.User, error) {
	// CreateUserParams構造体にデータをセット
	arg := sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: password,
		Username:     username,
		DisplayName:  sql.NullString{String: displayName, Valid: true}, // displayNameが空文字の場合も考慮
	}

	// ユーザーをDBに作成
	user, err := uc.dao.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	// 作成されたユーザー情報を返す
	return user, nil
}
