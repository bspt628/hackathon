package controller

import (
	"database/sql"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/user"
)

type UserController struct {
    userUsecase *usecase.UserUsecase // Usecaseへの依存性注入
}

func NewUserController(dbConn *sql.DB) *UserController {
	queries := sqlc.New(dbConn)
	userDAO := dao.NewUserDAO(queries)
	userUsecase := usecase.NewUserUsecase(userDAO)
	return &UserController{userUsecase: userUsecase}
}