package postdao

import (
	"context"
	"fmt"
)

func (dao *PostDAO) RestorePost(ctx context.Context, postID string) (bool, error) {
	// トランザクションの開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("failed to begin transaction")
	}
	defer tx.Rollback()
	
	result, err := dao.queries.RestorePost(ctx, postID)
	if err != nil {
		return false, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if affectedRows > 0 {
		// 返信元の投稿IDを取得
		replyToID, err := dao.queries.GetReplyToID(ctx, postID)
		if err != nil {
			return false, fmt.Errorf("返信元投稿の取得に失敗しました: %v", err)
		}

		// 返信元の投稿が存在し、返信元がNULLでない場合
		if replyToID.Valid {
			// 返信元のreplies_countをインクリメント
			err := dao.queries.IncrementReplyCount(ctx, replyToID.String)
			if err != nil {
				return false, fmt.Errorf("返信元投稿の返信数更新に失敗しました: %v", err)
			}
		}
	}

	return affectedRows > 0, nil
}