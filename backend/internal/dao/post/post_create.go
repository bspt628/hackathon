package postdao

import (
	"context"
	"hackathon/db/sqlc/generated"
)

func (dao *PostDAO) CreatePost(ctx context.Context, arg sqlc.CreatePostParams) error {
	return dao.queries.CreatePost(ctx, arg)
}

func (dao *PostDAO) CheckPostExists(ctx context.Context, postID string) (bool, error) {
	return dao.queries.CheckPostExists(ctx, postID)
}

func (dao *PostDAO) CheckRootPostValidity(ctx context.Context, rootPostID string) (bool, error) {
	return dao.queries.CheckRootPostValidity(ctx, rootPostID)
}

func (dao *PostDAO) IncrementReplyCount(ctx context.Context, postID string) error {
	return dao.queries.IncrementReplyCount(ctx, postID)
}