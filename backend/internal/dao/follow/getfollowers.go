package followdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
)

// GetFollowers retrieves followers of a given user by ID.
func (dao *FollowDAO) GetFollowers(ctx context.Context, userID string) ([]sqlc.GetFollowersRow, error) {
	userIDnull := sql.NullString{String: userID, Valid: true}
	followers, err := dao.queries.GetFollowers(ctx, userIDnull)
	if err != nil {
		return nil, err
	}
	return followers, nil
}