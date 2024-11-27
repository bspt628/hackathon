package postdao

import (
	"context"
	"hackathon/db/sqlc/generated"
)

func (dao *PostDAO) CreatePost(ctx context.Context, arg sqlc.CreatePostParams) error {
	return dao.db.CreatePost(ctx, arg)
}
