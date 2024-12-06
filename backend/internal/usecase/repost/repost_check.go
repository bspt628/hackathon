package repostusecase

import (
	"context"
	"errors"
	"hackathon/internal/model"
)

// GetRepostStatus はリポストのステータスを取得する
func (uc *RepostUsecase) GetRepostStatus(ctx context.Context, params model.GetRepostStatusParams) (bool, error) {
	// リポストのステータスを取得
	repost, err := uc.dao.GetRepostStatus(ctx, params)
	if err != nil {
		return false, errors.New("failed to check repost")
	}
	return repost, nil
}