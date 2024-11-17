package dao

import (
	"context"
	"github.com/oklog/ulid"
	"hackathon/db/sqlc/generated"
)


func (dao *UserDAO) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (*sqlc.User, error) {
	// IDをulidで自動生成する
	
	myid := ulid.MustNew(ulid.Now(), nil).String()
	// SQLクエリを実行して新しいユーザーを作成
	_, err := dao.db.CreateUser(ctx, sqlc.CreateUserParams{
		ID: 		  myid,
		Email:        arg.Email,
		PasswordHash: arg.PasswordHash,
		Username:     arg.Username,
		DisplayName:  arg.DisplayName,
	})
	if err != nil {
		return nil, err
	}

	// 新しく作成されたユーザーの ID で情報を再取得
	user, err := dao.GetUser(ctx, myid)
	if err != nil {
		return nil, err
	}

	return user, nil
}