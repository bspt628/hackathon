package followdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"fmt"
)

func (dao *FollowDAO) RemoveFollow(ctx context.Context,followerID, followingID string) error {
	// フォロー情報を保存
	result, err := dao.queries.RemoveFollow(ctx, sqlc.RemoveFollowParams{
		FollowerID: sql.NullString{String: followerID, Valid: true},
		FollowingID: sql.NullString{String: followingID, Valid: true},
	})

	// 削除された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no follow found from id %s to id %s", followerID, followingID)
	}
	return err
}