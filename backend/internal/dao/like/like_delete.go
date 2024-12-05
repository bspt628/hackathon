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

	// いいねを削除
	result, err := dao.queries.RemoveLike(ctx, arg)
	if err != nil {
		return errors.New("failed to delete like")
	}

	// 削除された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to check affected rows")
	}
	if rowsAffected == 0 {
		return errors.New("no like found")
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
