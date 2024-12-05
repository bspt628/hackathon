package repostdao

import (
	"context"
	"errors"
	"hackathon/internal/model"
)

func (dao *RepostDAO) CreateRepost(ctx context.Context, arg model.CreateRepostParams) error {
	// 引数を変換
	argdao := model.ConvertCreateRepostParamsToRepost(arg)

	// トランザクションを開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("failed to begin transaction")
	}
	defer tx.Rollback()

	// リポストを登録
	err = dao.queries.CreateRepost(ctx, argdao)
	if err != nil {
		return errors.New("failed to create repost")
	}

	// リポスト数をインクリメント
	if err := dao.queries.IncrementRepostsCount(ctx, arg.OriginalPostID); err != nil {
		return errors.New("failed to increment repost count")
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return errors.New("failed to commit transaction")
	}


	return nil
}

