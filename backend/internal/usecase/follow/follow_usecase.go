package followusecase

import (
	"hackathon/internal/dao/follow"
	"hackathon/internal/usecase/user"
)

type FollowUsecase struct {
	followDAO   *followdao.FollowDAO
	userUsecase *userusecase.UserUsecase
}

func NewFollowUsecase(followDAO *followdao.FollowDAO, userUsecase *userusecase.UserUsecase) *FollowUsecase {
	return &FollowUsecase{followDAO: followDAO, userUsecase: userUsecase}
}