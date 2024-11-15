package usecase

import (
	"hackathon/internal/dao"
)

type UserUsecase struct {
	dao *dao.UserDAO
}

func NewUserUsecase(dao *dao.UserDAO) *UserUsecase {
	return &UserUsecase{dao: dao}
}