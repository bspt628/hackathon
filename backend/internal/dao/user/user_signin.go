package dao

import (
	"context"
	"fmt"
	sqlc "hackathon/db/sqlc/generated"
)

type UserSignInDAO struct {
	db *sqlc.Queries
}

func NewUserSignInDAO(db *sqlc.Queries) *UserSignInDAO {
	return &UserSignInDAO{db: db}
}

// GetUserByEmail - メールアドレスからユーザー情報を取得
func (dao *UserSignInDAO) GetUserByEmail(ctx context.Context, email string) (string, string, error) {
	// GetUserByEmailを実行し、結果を取得
	user, err := dao.db.GetUserByEmail(ctx, email)
	if err != nil {
		return "", "", fmt.Errorf("ユーザー情報取得エラー: %v", err)
	}
	// userID と passwordHash を返す
	return user.ID, user.PasswordHash, nil
}