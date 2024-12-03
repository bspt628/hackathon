package followusecase

import (
	"context"
    "hackathon/db/sqlc/generated"
)

// GetFollowersUsecase retrieves followers using the DAO layer.
func (fc *FollowUsecase) GetFollowings(ctx context.Context, userID string) ([]sqlc.GetFollowingsRow, error) {
	return fc.followDAO.GetFollowings(ctx, userID)
}