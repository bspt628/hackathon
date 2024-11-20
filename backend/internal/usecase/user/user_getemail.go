package usecase

import (
	"context"
)

// ユースケース: usernameからemailを取得
func (uc *UserUsecase) GetEmailByUsernameUsecase(ctx context.Context, username string) (string, error) {
	return uc.dao.GetEmailByUsername(ctx, username)
}
