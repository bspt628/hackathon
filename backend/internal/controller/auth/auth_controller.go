package authController

import (
	"hackathon/internal/usecase/user"
	"hackathon/internal/dao/user"
	"database/sql"
)

type AuthController struct {
	signInUsecase *usecase.UserSignInUsecase
}



func NewAuthController(dbConn *sql.DB) *AuthController {
	signInDAO := dao.NewUserSignInDAO(dbConn)
	signInUsecase := usecase.NewUserSignInUsecase(signInDAO)
	return &AuthController{signInUsecase: signInUsecase}
}
