package likeusecase

import (
	"hackathon/internal/dao/like"
)

type LikeUsecase struct {
	dao *likedao.LikeDAO
}

func NewLikeUsecase(dao *likedao.LikeDAO) *LikeUsecase {
	return &LikeUsecase{dao: dao}
}