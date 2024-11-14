package usecase

import (
	"context"
	"hackathon/internal/dao"
	"hackathon/db/sqlc/generated" // ここでdb.User型をインポート
)

type UserUsecase struct {
	dao *dao.UserDAO
}

func NewUserUsecase(dao *dao.UserDAO) *UserUsecase {
	return &UserUsecase{dao: dao}
}

func (usecase *UserUsecase) GetUserByID(ctx context.Context, id string) (*db.User, error) {
	// dao層のGetUserByIDメソッドを呼び出す
	user, err := usecase.dao.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
