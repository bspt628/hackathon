package likedao

import (
	"context"
	"errors"
	"hackathon/db/sqlc/generated"
)

func (dao *LikeDAO) CreateLike(ctx context.Context, arg sqlc.AddLikeParams) error {
	// トランザクションを開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("failed to begin transaction")
	}
	defer tx.Rollback()

	// いいねを登録
	if err := dao.queries.AddLike(ctx, arg); err != nil {
		return errors.New("failed to add like")
	}

	// いいねのインクリメント
	if err := dao.queries.IncrementLikesCount(ctx, arg.PostID.String); err != nil {
		return errors.New("failed to increment like count")
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}