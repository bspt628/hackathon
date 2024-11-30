package userusecase

import (
	"context"
	"hackathon/db/sqlc/generated"
	"golang.org/x/crypto/bcrypt"
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func (uc *UserUsecase) CreateUser(ctx context.Context, email, password, username, displayname string) (*sqlc.User, error) {
	// IDをulidで自動生成する
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	myid := ulid.MustNew(ulid.Now(), entropy).String()

	// bcyptでパスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}


	// ユーザーをDBに作成
	user, err := uc.dao.CreateUser(ctx, myid, email, string(hashedPassword), username, displayname, password)
	if err != nil {
		return nil, err
	}

	// 作成されたユーザー情報を返す
	return user, nil
}
