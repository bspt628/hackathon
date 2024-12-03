package postusecase

import (
	"context"
	"hackathon/domain"
)

func (uc *PostUsecase) GetAllPosts(ctx context.Context, limit int32) (*domain.PostAll, error) {
    return uc.dao.GetAllPosts(ctx, limit)
}

func (uc *PostUsecase) GetFollowedUsersPosts(ctx context.Context, userID string, limit int32) (*domain.PostAll, error) {
    return uc.dao.GetFollowedUsersPosts(ctx, userID, limit)
}