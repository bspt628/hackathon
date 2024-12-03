package userdao

import (
	"context"
	"fmt"
	"hackathon/db/sqlc/generated"
	"time"
	"database/sql"

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
	if err := dao.queries.SaveResetToken(ctx, params); err != nil {
		return fmt.Errorf("トークン保存エラー: %w", err)
	}
	return nil
}


// ユーザーのパスワード更新
func (dao *UserDAO) UpdatePassword(ctx context.Context, tx *sql.Tx, email, newPassword string) error {
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("パスワードハッシュ化エラー: %w", err)
	}
	params := sqlc.UpdatePasswordParams{
		PasswordHash: string(hashedPassword),
		Email:        email,
	}

	if err := dao.queries.WithTx(tx).UpdatePassword(ctx, params); err != nil {
		return fmt.Errorf("パスワード更新エラー: %w", err)
	}
	return nil
}

// パスワードリセット
func (dao *UserDAO) ResetPassword(ctx context.Context, token, newPassword string) error {
	// トランザクションの開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// トークン検証
	email, err := dao.queries.WithTx(tx).ValidateResetToken(ctx, token)
	if err != nil {
		return fmt.Errorf("トークンが正しくありません: %w", err)
	}

	// パスワード更新
	if err := dao.UpdatePassword(ctx, tx, email, newPassword); err != nil {
		return fmt.Errorf("パスワード更新エラー: %w", err)
	}

	// トークン削除
	err = dao.queries.WithTx(tx).DeleteResetToken(ctx, token)
	if err != nil {
		return fmt.Errorf("トークン削除エラー: %w", err)
	}

	// トランザクションのコミット
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}