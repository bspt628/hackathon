package likeusecase

import (
	"context"
	"hackathon/db/sqlc/generated"
	"database/sql"
	"github.com/google/uuid"
)

func (lc *LikeUsecase) CreateLike(ctx context.Context, postID, userID string) error {
	// ulidを生成
	likeID := uuid.New().String()

	arg := sqlc.AddLikeParams{
		ID: likeID,
		UserID: sql.NullString{String: userID, Valid: true},
		PostID: sql.NullString{String: postID, Valid: true},
	}
	return lc.dao.CreateLike(ctx, arg)
}