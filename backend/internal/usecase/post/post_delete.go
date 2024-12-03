package postusecase

import (
	"context"
	"errors"
	"fmt"
)

func (pc *PostUsecase) DeletePost(ctx context.Context, postID string) error {
	if pc == nil || pc.dao == nil {
		return errors.New("PostUsecase or PostDAO is nil")
	}

	// 投稿の存在確認
	exists, err := pc.dao.CheckPostExists(ctx, postID)
	if err != nil {
		return fmt.Errorf("投稿の存在確認に失敗しました: %v", err)
	}
	if !exists {
		return fmt.Errorf("指定された post_id は存在しません")
	}

	// 投稿の削除
	err = pc.dao.DeletePost(ctx, postID)
	if err != nil {
		return fmt.Errorf("投稿の削除に失敗しました: %v", err)
	}

	return nil
}
