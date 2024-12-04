package followdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"github.com/go-sql-driver/mysql"
	"errors"
)

func (dao *FollowDAO) AddFollow(ctx context.Context, followID, followerID, followingID string) error {
	// フォロー情報を保存
	err := dao.queries.AddFollow(ctx, sqlc.AddFollowParams{
		ID: followID, 
		FollowerID: sql.NullString{String: followerID, Valid: true},
		FollowingID: sql.NullString{String: followingID, Valid: true},
	})

	if errmysql, ok := err.(*mysql.MySQLError); ok {
		if errmysql.Number == 1062 {
			return errors.New("follow already exists")
		}
	}else if err != nil {
		return errors.New("failed to add follow")
	}
	return nil
}