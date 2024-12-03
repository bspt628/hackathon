package likedao

import (
	"context"
	"errors"
	"fmt"
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
	result, err := dao.queries.AddLike(ctx, arg)
	if err != nil {
		return errors.New("failed to add like")
	}
	// いいねがすでに存在する場合
	count, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to check affected rows")
	}
	fmt.Println("count: ", count)
	if count == 0 {
		return errors.New("like already exists")
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