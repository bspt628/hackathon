package followdao

import (
	"context"
	"hackathon/db/sqlc/generated"
	"database/sql"
)

// GetFollowers retrieves followers of a given user by ID.
func (dao *FollowDAO) GetFollowings(ctx context.Context, userID string) ([]sqlc.GetFollowingsRow, error) {
	userIDnull := sql.NullString{String: userID, Valid: true}
	followers, err := dao.queries.GetFollowings(ctx, userIDnull)
	if err != nil {
		return nil, err
	}
	return followers, nil
}