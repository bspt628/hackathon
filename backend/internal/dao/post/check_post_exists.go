package postdao

import (
	"context"
)

func (dao *PostDAO) CheckPostExists(ctx context.Context, postID string) (bool, error) {
	return dao.queries.CheckPostExists(ctx, postID)
}

func (dao *PostDAO) CheckRootPostValidity(ctx context.Context, rootPostID string) (bool, error) {
	return dao.queries.CheckRootPostValidity(ctx, rootPostID)
}
