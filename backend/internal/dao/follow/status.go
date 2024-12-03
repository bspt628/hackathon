package followdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
)

func (dao *FollowDAO) GetFollowStatus(ctx context.Context, followerID, followingID string) (bool, error) {
	// フォロー情報を保存
	return dao.queries.GetFollowStatus(ctx, sqlc.GetFollowStatusParams{
		FollowerID: sql.NullString{String: followerID, Valid: true},
		FollowingID: sql.NullString{String: followingID, Valid: true},
	})
	
}