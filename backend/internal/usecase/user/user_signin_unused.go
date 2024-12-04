package userusecase

import (
	"context"
	"errors"
	"hackathon/internal/auth/jwt_unused"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// サインイン処理
func (uc *UserUsecase) SignIn(ctx context.Context, username, password string) (string, error) {
	// DAOからユーザー情報を取得
	// emailを出力
	userID, passwordHash, err := uc.dao.GetUserPasswordFromUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// パスワードの検証
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return "", errors.New("パスワードが一致しません")
	}

	// トークンを生成
	token, err := jwtunused.GenerateToken(userID, time.Hour*24)
	if err != nil {
		return "", err
	}

	return token, nil
}
