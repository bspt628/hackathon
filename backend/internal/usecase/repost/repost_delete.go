package repostusecase

import (
	"context"
	"hackathon/internal/model"
)

func (rc *RepostUsecase) DeleteRepost(ctx context.Context, params model.DeleteRepostParams) error {
	// リポスト情報を削除
	return rc.dao.DeleteRepost(ctx, params)
}
