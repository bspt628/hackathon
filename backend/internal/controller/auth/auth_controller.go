package authcontroller

import (
	"database/sql"
	"hackathon/internal/usecase/user"
	"hackathon/internal/dao/user"
	"hackathon/db/sqlc/generated"
)

type AuthController struct {
	signInUsecase *userusecase.UserSignInUsecase
}

func NewAuthController(dbConn *sql.DB) *AuthController {
	queries := sqlc.New(dbConn)
	signInDAO := userdao.NewUserSignInDAO(queries)
	signInUsecase := userusecase.NewUserSignInUsecase(signInDAO)
	return &AuthController{signInUsecase: signInUsecase}
}
