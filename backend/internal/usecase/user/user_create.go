package userusecase

import (
	"context"
	"hackathon/db/sqlc/generated"
	"golang.org/x/crypto/bcrypt"
	"github.com/oklog/ulid"
	"math/rand"
	"time"
	"strings"
	"errors"
)

func (uc *UserUsecase) CreateUser(ctx context.Context, email, password, username, displayname string) (*sqlc.User, error) {
	// IDをulidで自動生成する
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	myid := ulid.MustNew(ulid.Now(), entropy).String()

	// usernameに特殊文字が含まれているかチェック
	if !uc.CheckUsername(username) {
		return nil, errors.New("usernameに特殊文字が含まれています")
	}

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

// CheckUsername はusernameに特殊文字が含まれているかチェックする
func (uc *UserUsecase) CheckUsername(username string) bool {
	// usernameに特殊文字が含まれているかチェック
	return  !strings.ContainsAny(username, "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~") 
}