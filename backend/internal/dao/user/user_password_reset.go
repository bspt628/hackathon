package userdao

import (
	"context"
	"fmt"
	"hackathon/db/sqlc/generated"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// トークンの保存
func (dao *UserDAO) SaveResetToken(ctx context.Context, email, token string, expiry time.Time) error {
	// SaveResetTokenParams にデータを格納
	params := sqlc.SaveResetTokenParams{
		Email:  email,
		Token:  token,
		Expiry: expiry,
	}

	// 生成された SaveResetToken メソッドを呼び出す
	if err := dao.db.SaveResetToken(ctx, params); err != nil {
		return fmt.Errorf("トークン保存エラー: %w", err)
	}
	return nil
}

// トークンの検証
func (dao *UserDAO) ValidateResetToken(ctx context.Context, token string) (string, error) {
	email, err := dao.db.ValidateResetToken(ctx, token)
	if err != nil {
		return "", fmt.Errorf("トークン検証エラー: %w", err)
	}
	return email, nil
}

// トークンの削除
func (dao *UserDAO) DeleteResetToken(ctx context.Context, token string) error {
	if err := dao.db.DeleteResetToken(ctx, token); err != nil {
		return fmt.Errorf("トークン削除エラー: %w", err)
	}
	return nil
}

// ユーザーのパスワード更新
func (dao *UserDAO) UpdatePasswordByEmail(ctx context.Context, email, newPassword string) error {
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("パスワードハッシュ化エラー: %w", err)
	}
	params := sqlc.UpdatePasswordByEmailParams{
		PasswordHash: string(hashedPassword),
		Email:        email,
	}

	if err := dao.db.UpdatePasswordByEmail(ctx, params); err != nil {
		return fmt.Errorf("パスワード更新エラー: %w", err)
	}
	return nil
}
