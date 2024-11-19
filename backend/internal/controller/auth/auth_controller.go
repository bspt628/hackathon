package authController

import (
	"hackathon/internal/usecase/user"
	"hackathon/internal/dao/user"
	"hackathon/internal/email"
	"database/sql"
)

type AuthController struct {
	signInUsecase *usecase.UserSignInUsecase
}

type PasswordResetController struct {
	passwordResetUsecase *usecase.UserPasswordResetUsecase
}

func NewAuthController(dbConn *sql.DB) *AuthController {
	signInDAO := dao.NewUserSignInDAO(dbConn)
	signInUsecase := usecase.NewUserSignInUsecase(signInDAO)
	return &AuthController{signInUsecase: signInUsecase}
}


func NewPasswordResetController(dbConn *sql.DB) *PasswordResetController {
	passwordResetDAO := dao.NewUserPasswordResetDAO(dbConn)
	emailSender := email.NewEmailSender("smtp.gmail.com", "587", "your-email@gmail.com", "your-app-password")
	passwordResetUsecase := usecase.NewUserPasswordResetUsecase(passwordResetDAO, emailSender)
	return &PasswordResetController{passwordResetUsecase: passwordResetUsecase}
}
