package postdao

import (
	"context"
	"fmt"
)

func (dao *PostDAO) DeletePost(ctx context.Context, postID string) error {
	// トランザクションの開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("failed to begin transaction")
	}
	defer tx.Rollback()

	// 返信元のIDを取得
	replyToID, err := dao.queries.GetReplyToID(ctx, postID)
	if err != nil {
		return fmt.Errorf("failed to get reply_to_id: %v", err)
	}

	// 投稿の削除
	err = dao.queries.DeletePost(ctx, postID)
	if err != nil {
		return fmt.Errorf("failed to delete post: %v", err)
	}

	// 返信元の reply_to_id を 出力
	fmt.Println("reply_to_id: ", replyToID)

	// 返信元の reply_count をデクリメント
	if replyToID.Valid {
		err = dao.queries.DecrementReplyCount(ctx, replyToID.String)
		if err != nil {
			return fmt.Errorf("failed to decrement reply_count: %v", err)
		}
	}

	// トランザクションのコミット
	err = tx.Commit()
	if err != nil {
		fmt.Println("failed to commit transaction")
	}

	return nil
}

