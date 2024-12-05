package followdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"fmt"
)

func (dao *FollowDAO) RemoveFollow(ctx context.Context, followerID, followingID string) error {
	// トランザクションを開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// フォロー情報を保存
	result, err := dao.queries.RemoveFollow(ctx, sqlc.RemoveFollowParams{
		FollowerID: sql.NullString{String: followerID, Valid: true},
		FollowingID: sql.NullString{String: followingID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to remove follow: %w", err)
	}

	// 削除された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no follow found from id %s to id %s", followerID, followingID)
	}

	// フォロワー数をデクリメント
	_, err = dao.queries.DecrementFollowersCount(ctx, followingID)
	if err != nil {
		return fmt.Errorf("failed to decrement followers count: %w", err)
	}

	// フォロー数をデクリメント
	_, err = dao.queries.DecrementFollowingsCount(ctx, followerID)
	if err != nil {
		return fmt.Errorf("failed to decrement followings count: %w", err)
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return err
}