package dao

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type UserPasswordResetDAO struct {
	db *sql.DB
}

func NewUserPasswordResetDAO(db *sql.DB) *UserPasswordResetDAO {
	return &UserPasswordResetDAO{db: db}
}

// トークンの保存
func (dao *UserPasswordResetDAO) SaveResetToken(ctx context.Context, email, token string, expiry time.Time) error {
	_, err := dao.db.ExecContext(ctx, "INSERT INTO password_reset_tokens (email, token, expiry) VALUES (?, ?, ?)", email, token, expiry)
	return err
}

// トークンの検証
func (dao *UserPasswordResetDAO) ValidateResetToken(ctx context.Context, token string) (string, error) {
	var email string
	err := dao.db.QueryRowContext(ctx, "SELECT email FROM password_reset_tokens WHERE token = ? AND expiry > NOW()", token).Scan(&email)
	if err != nil {
		return "", fmt.Errorf("トークン検証エラー: %v", err)
	}
	return email, nil
}

// トークンの削除
func (dao *UserPasswordResetDAO) DeleteResetToken(ctx context.Context, token string) error {
	_, err := dao.db.ExecContext(ctx, "DELETE FROM password_reset_tokens WHERE token = ?", token)
	return err
}

// ユーザーのパスワード更新
func (dao *UserPasswordResetDAO) UpdatePasswordByEmail(ctx context.Context, email, newPassword string) error {
	// パスワードをハッシュ化する処理（ここでは仮に新しいパスワードをそのまま使用）
	// 実際にはハッシュ化して保存するべきです。
	_, err := dao.db.ExecContext(ctx, "UPDATE users SET password = ? WHERE email = ?", newPassword, email)
	if err != nil {
		return fmt.Errorf("パスワード更新エラー: %v", err)
	}
	return nil
}
