package postusecase

import (
	"context"
	"errors"
)

func (uc *PostUsecase) RestorePost(ctx context.Context, postID string) error {
	// DAO層に処理を委譲
	success, err := uc.dao.RestorePost(ctx, postID)
	if err != nil {
		return err
	}

	// 復活できなかった場合のエラーハンドリング
	if !success {
		return errors.New("投稿の復活に失敗しました（削除から10分以上経過している可能性があります）")
	}

	return nil
}
