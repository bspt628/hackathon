package followdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
)

func (dao *FollowDAO) AddFollow(ctx context.Context,followID, followerID, followeeID string) error {
	// フォロー情報を保存
	err := dao.db.AddFollow(ctx, sqlc.AddFollowParams{
		ID: followID, 
		FollowerID: sql.NullString{String: followerID, Valid: true},
		FollowingID: sql.NullString{String: followeeID, Valid: true},
	})
	return err
}