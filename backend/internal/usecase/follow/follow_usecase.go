package followusecase

import (
	"hackathon/internal/dao/follow"
	"hackathon/internal/usecase/user"
)

type FollowUsecase struct {
	followDAO   *followdao.FollowDAO
	userUsecase *usecase.UserUsecase
}

func NewFollowUsecase(followDAO *followdao.FollowDAO, userUsecase *usecase.UserUsecase) *FollowUsecase {
	return &FollowUsecase{followDAO: followDAO, userUsecase: userUsecase}
}