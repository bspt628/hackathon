package repostdao

import (
	"context"
	"hackathon/internal/model"
	"errors"
)

// DeleteRepost はリポスト情報を削除する
func (dao *RepostDAO) DeleteRepost(ctx context.Context, params model.DeleteRepostParams) error {
	// sqlへ型変換
	argdao := model.ConvertDeleteRepostParamsToRepost(params)

	// トランザクションの開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("failed to begin transaction")
	}
	defer tx.Rollback()

	// リポストを削除
	result, err := dao.queries.DeleteRepost(ctx, argdao)
	if err != nil {
		return errors.New("failed to delete repost")
	}
	// 削除された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to check affected rows")
	}
	if rowsAffected == 0 {
		return errors.New("no repost found")
	}

	// リポスト数をデクリメント
	if err := dao.queries.DecrementRepostsCount(ctx, params.OriginalPostID); err != nil {
		return errors.New("failed to decrement repost count")
	}

	// トランザクションのコミット
	err = tx.Commit()
	if err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}