package followusecase

import (
	"context"
	"hackathon/db/sqlc/generated"
)

// GetFollowersUsecase retrieves followers using the DAO layer.
func (fc *FollowUsecase) GetFollowers(ctx context.Context, userID string) ([]sqlc.GetFollowersRow, error) {
	return fc.followDAO.GetFollowers(ctx, userID)
}