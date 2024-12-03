package userdao

import (
	"context"
	"hackathon/db/sqlc/generated"
)

func (dao *UserDAO) GetUser(ctx context.Context, id string) (*sqlc.User, error) {
	// GetUserByIdを呼び出して、返り値のGetUserByIdRowを受け取る
	row, err := dao.queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	// GetUserByIdRowからsqlc.Userに変換
	user := &sqlc.User{
		ID:             row.ID,
		FirebaseUid:    row.FirebaseUid,
		Email:          row.Email,
		Username:       row.Username,
		DisplayName:    row.DisplayName,
		Bio:            row.Bio,
		Location:       row.Location,
		FollowersCount: row.FollowersCount,
		FollowingCount: row.FollowingCount,
		PostsCount:     row.PostsCount,
	}

	return user, nil
}
