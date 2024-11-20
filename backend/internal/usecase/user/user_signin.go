package usecase

import (
	"context"
	"errors"
	"fmt"
	"hackathon/internal/auth"

	"time"

	"golang.org/x/crypto/bcrypt"
)

// サインイン処理
func (uc *UserSignInUsecase) SignIn(ctx context.Context, username, password string) (string, error) {
	// DAOからユーザー情報を取得
	// emailを出力
	fmt.Printf("username: %s\n", username)
	userID, passwordHash, err := uc.dao.GetUserByEmail(ctx, username)
	if err != nil {
		return "", err
	}
	fmt.Printf("userID: %s, passwordHash: [%s] [%s]\n", userID, passwordHash, password)

	// パスワードの検証
	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return "", errors.New("パスワードが一致しません")
	}

	// トークンを生成
	token, err := auth.GenerateToken(userID, time.Hour*24)
	if err != nil {
		return "", err
	}

	return token, nil
}
