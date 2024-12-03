package followusecase

import (
	"hackathon/internal/dao/follow"
)

type FollowUsecase struct {
	followDAO   *followdao.FollowDAO
}

func NewFollowUsecase(followDAO *followdao.FollowDAO) *FollowUsecase {
	return &FollowUsecase{
		followDAO: followDAO,
	}
}