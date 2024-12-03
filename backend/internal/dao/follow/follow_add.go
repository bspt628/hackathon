package followdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
)

func (dao *FollowDAO) AddFollow(ctx context.Context, followID, followerID, followingID string) error {
	// フォロー情報を保存
	return dao.queries.AddFollow(ctx, sqlc.AddFollowParams{
		ID: followID, 
		FollowerID: sql.NullString{String: followerID, Valid: true},
		FollowingID: sql.NullString{String: followingID, Valid: true},
	})
}