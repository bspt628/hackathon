package likeusecase

import (
	"context"
)

func (lc LikeUsecase) GetPostLikesCount(ctx context.Context, postID string) (int32, error) {
	return lc.dao.GetPostLikesCount(ctx, postID)
}