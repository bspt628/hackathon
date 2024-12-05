package likeusecase

import (
	"context"
	"database/sql"
	sqlc "hackathon/db/sqlc/generated"
)

func (lc *LikeUsecase) GetLikeStatus(ctx context.Context, userID, postID string) (bool, error) {
	arg := sqlc.GetLikeStatusParams{
		UserID: sql.NullString{String: userID, Valid: true},
		PostID: sql.NullString{String: postID, Valid: true},
	}
	like, err := lc.dao.GetLikeStatus(ctx, arg)
	if err != nil {
		return false, err
	}
	return like, nil
}
