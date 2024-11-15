package dao

import (
	"context"
	"github.com/oklog/ulid"
	"hackathon/db/sqlc/generated"
)

type UserDAO struct {
	db *db.Queries
}

func NewUserDAO(db *db.Queries) *UserDAO {
	return &UserDAO{db: db}
}

func (dao *UserDAO) GetUserByID(ctx context.Context, id string) (*db.User, error) {
	// GetUserByIdを呼び出して、返り値のGetUserByIdRowを受け取る
	row, err := dao.db.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	// GetUserByIdRowからdb.Userに変換
	user := &db.User{
		ID:            row.ID,
		Email:         row.Email,
		Username:      row.Username,
		DisplayName:   row.DisplayName,
		Bio:           row.Bio,
		Location:      row.Location,
		FollowersCount: row.FollowersCount,
		FollowingCount: row.FollowingCount,
		PostsCount:    row.PostsCount,
	}

	return user, nil
}

func (dao *UserDAO) CreateUser(ctx context.Context, arg db.CreateUserParams) (*db.User, error) {
	// IDをulidで自動生成する
	
	myid := ulid.MustNew(ulid.Now(), nil).String()
	// SQLクエリを実行して新しいユーザーを作成
	_, err := dao.db.CreateUser(ctx, db.CreateUserParams{
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
	user, err := dao.GetUserByID(ctx, myid)
	if err != nil {
		return nil, err
	}

	return user, nil
}