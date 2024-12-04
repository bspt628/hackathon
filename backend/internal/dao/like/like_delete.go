package likedao

import (
	"context"
	"errors"
	sqlc "hackathon/db/sqlc/generated"
)

func (dao *LikeDAO) DeleteLike(ctx context.Context, arg sqlc.RemoveLikeParams) error {
	// トランザクションを開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("failed to begin transaction")
	}
	defer tx.Rollback()

	argcheck := sqlc.CheckLikeExistsParams(arg)

	// いいねが存在するか確認
	exists, err := dao.queries.CheckLikeExists(ctx, argcheck)
	if err != nil {
		return errors.New("failed to check like exists")
	}
	if !exists {
		return errors.New("like does not exist")
	}

	// いいねを削除
	if err := dao.queries.RemoveLike(ctx, arg); err != nil {
		return errors.New("failed to delete like")
	}

	// いいねのデクリメント
	if err := dao.queries.DecrementLikesCount(ctx, arg.PostID.String); err != nil {
		return errors.New("failed to decrement like count")
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}
