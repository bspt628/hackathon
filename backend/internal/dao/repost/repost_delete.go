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

	// リポストが存在するか確認
	argcheck := model.ConvertGetRepostStatusParamsToRepost(params)
	exists, err := dao.queries.GetRepostStatus(ctx, argcheck)
	if err != nil {
		return errors.New("failed to check repost exists")
	}
	if !exists {
		return errors.New("repost does not exist")
	}

	// リポストを削除
	if err := dao.queries.DeleteRepost(ctx, argdao); err != nil {
		return errors.New("failed to delete repost")
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