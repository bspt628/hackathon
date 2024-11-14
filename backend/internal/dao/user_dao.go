package dao

import (
	"context"
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
