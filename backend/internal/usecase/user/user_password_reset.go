package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)



// パスワードリセットリクエスト
func (u *UserPasswordResetUsecase) RequestPasswordReset(ctx context.Context, email string) error {
	// トークン生成
	tokenBytes := make([]byte, 32)
	if _, err := rand.Read(tokenBytes); err != nil {
		return fmt.Errorf("トークン生成エラー: %v", err)
	}
	token := base64.URLEncoding.EncodeToString(tokenBytes)

	// トークンの保存
	expiry := time.Now().Add(1 * time.Hour)
	if err := u.passwordResetDAO.SaveResetToken(ctx, email, token, expiry); err != nil {
		return fmt.Errorf("トークン保存エラー: %v", err)
	}

	// メール送信
	resetLink := fmt.Sprintf("https://example.com/reset-password?token=%s", token)
	body := fmt.Sprintf("以下のリンクをクリックしてパスワードをリセットしてください:\n%s", resetLink)
	return u.emailSender.SendEmail(email, "パスワードリセット", body)
}

// パスワードリセット処理
func (u *UserPasswordResetUsecase) ResetPassword(ctx context.Context, token, newPassword string) error {
	// トークン検証
	email, err := u.passwordResetDAO.ValidateResetToken(ctx, token)
	if err != nil {
		return fmt.Errorf("トークン検証エラー: %v", err)
	}

	// パスワードの更新
	// ここでは仮にパスワードを更新する関数 UpdatePasswordByEmail を使用します。
	if err := u.passwordResetDAO.UpdatePasswordByEmail(ctx, email, newPassword); err != nil {
		return fmt.Errorf("パスワード更新エラー: %v", err)
	}

	// トークン削除
	return u.passwordResetDAO.DeleteResetToken(ctx, token)
}
