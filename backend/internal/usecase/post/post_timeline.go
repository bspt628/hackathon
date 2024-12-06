package postusecase

import (
	"context"
	"hackathon/internal/model"
)

func (uc *PostUsecase) GetAllPosts(ctx context.Context, limit int32) ([]model.PostAll, error) {
    return uc.dao.GetAllPosts(ctx, limit)
}

func (uc *PostUsecase) GetFollowingUsersPosts(ctx context.Context, userID string, limit int32) ([]model.PostAll, error) {
    return uc.dao.GetFollowingUsersPosts(ctx, userID, limit)
}