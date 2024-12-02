package postdao

import (
	"context"
)

func (dao *PostDAO) IncrementReplyCount(ctx context.Context, postID string) error {
	return dao.queries.IncrementReplyCount(ctx, postID)
}