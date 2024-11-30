package userusecase

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"golang.org/x/crypto/bcrypt"
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func (uc *UserUsecase) CreateUser(ctx context.Context, email, password, username, displayName string) (*sqlc.User, error) {
	// IDをulidで自動生成する
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	myid := ulid.MustNew(ulid.Now(), entropy).String()

	// bcyptでパスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// CreateUserParams構造体にデータをセット
	arg := sqlc.CreateUserParams{
		ID: 		  myid,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Username:     username,
		DisplayName:  sql.NullString{String: displayName, Valid: true}, // displayNameが空文字の場合も考慮
	}

	// ユーザーをDBに作成
	user, err := uc.dao.CreateUser(ctx, arg, password)
	if err != nil {
		return nil, err
	}

	// 作成されたユーザー情報を返す
	return user, nil
}
