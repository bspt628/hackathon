package repostdao

import (
	"context"
	"database/sql"
	"errors"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/model"
)

func (dao *RepostDAO) CreateRepost(ctx context.Context, arg model.CreateRepostParams) error {
	argdao := sqlc.CreateRepostParams{
		ID: arg.ID,
		UserID: sql.NullString{String: arg.UserID, Valid: true},
		OriginalPostID: sql.NullString{String: arg.OriginalPostID, Valid: true},
		IsQuoteRepost: sql.NullBool{Bool: arg.IsQuoteRepost, Valid: true},
		AdditionalComment: sql.NullString{String: arg.AdditionalComment, Valid: true},
	}

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
	_, err = dao.queries.IncrementRepostsCount(ctx, arg.OriginalPostID)
	if err != nil {
		return errors.New("failed to increment repost count")
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return errors.New("failed to commit transaction")
	}


	return nil
}

