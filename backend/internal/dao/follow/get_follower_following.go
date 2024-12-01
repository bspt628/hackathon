
package followdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
)

// GetFollowers retrieves followers of a given user by ID.
func (dao *FollowDAO) GetFollowersAndFollowings(ctx context.Context, followerID, followingID string) ([]sqlc.GetFollowersAndFollowingsRow, error) {
	followingIDnull := sql.NullString{String: followingID, Valid: true}
	followerIDnull := sql.NullString{String: followerID, Valid: true}
	FF, err := dao.queries.GetFollowersAndFollowings(ctx, sqlc.GetFollowersAndFollowingsParams{
		FollowingID: followingIDnull,
		FollowerID: followerIDnull,
	})
	if err != nil {
		return nil, err
	}
	return FF, nil
}