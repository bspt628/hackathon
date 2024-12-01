package followusecase

import (
	"context"
	"hackathon/db/sqlc/generated"
)

// GetFollowersUsecase retrieves followers using the DAO layer.
func (fc *FollowUsecase) GetFollowersAndFollowings(ctx context.Context, followerID, followingID string) ([]sqlc.GetFollowersAndFollowingsRow, error) {
	return fc.followDAO.GetFollowersAndFollowings(ctx, followerID, followingID)
}