package repostdao

import (
	"context"
	"hackathon/internal/model"
	"errors"
)

func (dao *RepostDAO) GetRepostStatus(ctx context.Context, params model.GetRepostStatusParams) (bool, error) {
	// sqlへ型変換
	argdao := model.ConvertGetRepostStatusParamsToRepost(params)

	// リポストのステータスを取得
	repost, err := dao.queries.GetRepostStatus(ctx, argdao)
	if err != nil {
		return false, errors.New("failed to check repost")
	}
	return repost, nil
}