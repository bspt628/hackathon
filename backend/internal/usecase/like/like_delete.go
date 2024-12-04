package likeusecase

import (
	"context"
	"database/sql"
	sqlc "hackathon/db/sqlc/generated"
)

func (lc *LikeUsecase) DeleteLike(ctx context.Context, postID, userID string) error {
	// いいね情報を削除
	return lc.dao.DeleteLike(ctx, sqlc.RemoveLikeParams{
		PostID: sql.NullString{String: postID, Valid: true},
		UserID: sql.NullString{String: userID, Valid: true},
	})
}