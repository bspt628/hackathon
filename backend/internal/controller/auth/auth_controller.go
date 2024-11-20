package authController

import (
	"database/sql"
	"hackathon/internal/usecase/user"
	"hackathon/internal/dao/user"
	"hackathon/db/sqlc/generated"
)

type AuthController struct {
	signInUsecase *usecase.UserSignInUsecase
}



func NewAuthController(dbConn *sql.DB) *AuthController {
	queries := sqlc.New(dbConn)
	signInDAO := dao.NewUserSignInDAO(queries)
	signInUsecase := usecase.NewUserSignInUsecase(signInDAO)
	return &AuthController{signInUsecase: signInUsecase}
}
