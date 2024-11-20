package usecase

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
)

func (uc *UserUsecase) CreateUser(ctx context.Context, email, passwordHash, username, displayName string) (*sqlc.User, error) {
	// CreateUserParams構造体にデータをセット
	arg := sqlc.CreateUserParams{
		Email:        email,
		PasswordHash: passwordHash,
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
