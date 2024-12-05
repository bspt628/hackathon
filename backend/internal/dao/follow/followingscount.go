package followdao

import (
	"context"
	"fmt"
)

func (dao *FollowDAO) GetFollowingsCount(ctx context.Context, userID string) (int32, error) {

	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	// トランザクション用クエリオブジェクトの作成
	qtx := dao.queries.WithTx(tx)

	// 最新のフォロワー数を取得
	nullCount, err := qtx.GetFollowingsCount(ctx, userID)
	if err != nil {
		tx.Rollback() // ロールバック
		return 0, fmt.Errorf("failed to get following count: %w", err)
	}

	// トランザクションのコミット
	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// sql.NullInt32 を int32 に変換
	if !nullCount.Valid {
		return 0, fmt.Errorf("followers count is NULL for user ID: %s", userID)
	}

	return nullCount.Int32, nil

}