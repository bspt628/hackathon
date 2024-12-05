package repostusecase

import (
	"hackathon/internal/dao/repost"
)

type RepostUsecase struct {
	dao *repostdao.RepostDAO
}

func NewRepostUsecase(dao *repostdao.RepostDAO) *RepostUsecase {
	return &RepostUsecase{dao: dao}
}
