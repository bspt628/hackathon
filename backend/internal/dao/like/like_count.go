package likedao

import(
	"context"
	"errors"
)

func (dao *LikeDAO) GetPostLikesCount(ctx context.Context, postID string) (int32, error) {
	count, err := dao.queries.GetPostLikesCount(ctx, postID)
	if err != nil {
		return 0, errors.New("failed to get post likes count")
	}
	return count.Int32, nil
}