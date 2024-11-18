package dao

import (
	"context"
	"database/sql"
	"errors"
)

type UserSignInDAO struct {
	db *sql.DB
}

func NewUserSignInDAO(db *sql.DB) *UserSignInDAO {
	return &UserSignInDAO{db: db}
}

// メールアドレスからユーザー情報を取得
func (dao *UserSignInDAO) GetUserByEmail(ctx context.Context, email string) (string, string, error) {
	var userID string
	var passwordHash string

	query := "SELECT id, password_hash FROM users WHERE email = ?"
	err := dao.db.QueryRowContext(ctx, query, email).Scan(&userID, &passwordHash)
	if err == sql.ErrNoRows {
		return "", "", errors.New("ユーザーが存在しません")
	}
	if err != nil {
		return "", "", err
	}

	return userID, passwordHash, nil
}
