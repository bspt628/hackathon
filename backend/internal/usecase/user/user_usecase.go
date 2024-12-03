package userusecase

import (
	"hackathon/internal/dao/user"
)

type UserUsecase struct {
	dao *userdao.UserDAO
}

func NewUserUsecase(dao *userdao.UserDAO) *UserUsecase {
	return &UserUsecase{dao: dao}
}