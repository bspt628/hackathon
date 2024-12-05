package repostusecase

import (
	"context"
	"hackathon/internal/model"
)

func (rc *RepostUsecase) CreateRepost(ctx context.Context, params model.CreateRepostParams) error {
	
	// リポスト情報をDAOを通じて保存
	return rc.dao.CreateRepost(ctx, params)
}