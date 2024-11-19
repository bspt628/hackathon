package usecase

import (
	"hackathon/internal/dao/user"
	"hackathon/internal/email"
)

type UserUsecase struct {
	dao *dao.UserDAO
}

func NewUserUsecase(dao *dao.UserDAO) *UserUsecase {
	return &UserUsecase{dao: dao}
}

type UserSignInUsecase struct {
	dao *dao.UserSignInDAO
}

func NewUserSignInUsecase(dao *dao.UserSignInDAO) *UserSignInUsecase {
	return &UserSignInUsecase{dao: dao}
}


type UserPasswordResetUsecase struct {
	passwordResetDAO *dao.UserPasswordResetDAO
	emailSender *email.EmailSender
}

func NewUserPasswordResetUsecase(passwordResetDAO *dao.UserPasswordResetDAO, emailSender *email.EmailSender) *UserPasswordResetUsecase {
	return &UserPasswordResetUsecase{
		passwordResetDAO: passwordResetDAO,
	}
}