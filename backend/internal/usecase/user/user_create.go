package userusecase

import (
	"context"
	"errors"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/model"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUsecase) CreateUser(ctx context.Context, request model.UserCreateRequest) (*sqlc.User, error) {

	// usernameに特殊文字が含まれているかチェック
	if !uc.CheckUsername(request.Username) {
		return nil, errors.New("usernameに特殊文字が含まれています")
	}

	// bcyptでパスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// DAO層に送るリクエストを作成
	requestDAO := model.NewUserCreateDAORequest(request.Email, string(hashedPassword), request.Username, request.DisplayName)


	// ユーザーをDBに作成
	user, err := uc.dao.CreateUser(ctx, requestDAO)
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