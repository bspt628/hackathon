package userusecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
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
	expiry := time.Now().Add(10 * time.Hour) // 世界標準時だった！
	if err := u.passwordResetDAO.SaveResetToken(ctx, email, token, expiry); err != nil {
		return fmt.Errorf("トークン保存エラー: %v", err)
	}

	// メール送信
	resetLink := fmt.Sprintf("http://localhost:8080/reset-password?token=%s", token)
	body := fmt.Sprintf("以下のリンクをクリックしてパスワードをリセットしてください:\n\n%s", resetLink)
	subject := "パスワードリセットのお知らせ"

	// メール送信処理
	err := sendEmail(email, subject, body)
	if err != nil {
		return fmt.Errorf("メール送信エラー: %v", err)
	}

	return nil
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

// メール送信
func sendEmail(to, subject, body string) error {
	// メールのヘッダーとボディを組み立てる
	// SMTP設定
	SMTPHost := os.Getenv("SMTP_HOST")
	SMTPPort := os.Getenv("SMTP_PORT")
	FromEmail := os.Getenv("FROM_EMAIL") // 送信元メールアドレス

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", FromEmail, to, subject, body)

	// SMTPサーバーに接続してメールを送信
	addr := fmt.Sprintf("%s:%s", SMTPHost, SMTPPort)
	err := smtp.SendMail(addr, nil, FromEmail, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("SMTP送信エラー: %v", err)
	}
	return nil
}
