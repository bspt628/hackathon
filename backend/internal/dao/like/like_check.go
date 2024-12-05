package likedao

import (
	"context"
	"errors"
	"hackathon/db/sqlc/generated"
)

func (dao *LikeDAO) GetLikeStatus(ctx context.Context, arg sqlc.GetLikeStatusParams) (bool, error) {
	like, err := dao.queries.GetLikeStatus(ctx, arg)
	if err != nil {
		return false, errors.New("failed to check like")
	}
	return like, nil
}