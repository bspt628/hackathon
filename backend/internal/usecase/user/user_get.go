package userusecase

import (
	"context"
	"hackathon/db/sqlc/generated"
)

// ユースケース: usernameからemailを取得
func (uc *UserUsecase) GetEmailByUsernameUsecase(ctx context.Context, username string) (string, error) {
	email, err := uc.dao.GetEmailByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	return email, nil
}

// ユースケース:
func (uc *UserUsecase) GetUser(ctx context.Context, id string) (*sqlc.User, error) {
	user, err := uc.dao.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
