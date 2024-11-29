package userdao

import (
	"context"
	"fmt"
)

// GetUserByEmail - メールアドレスからユーザー情報を取得
func (dao *UserDAO) GetUserByEmail(ctx context.Context, email string) (string, string, error) {
	// GetUserByEmailを実行し、結果を取得
	user, err := dao.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", fmt.Errorf("ユーザー情報取得エラー: %v", err)
	}
	// userID と passwordHash を返す
	return user.ID, user.PasswordHash, nil
}
