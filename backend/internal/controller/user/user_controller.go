package usercontroller

import (
	"database/sql"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/user"
)

type UserController struct {
    userUsecase *userusecase.UserUsecase // Usecaseへの依存性注入
}

func NewUserController(dbConn *sql.DB) *UserController {
	queries := sqlc.New(dbConn)
	userDAO := userdao.NewUserDAO(queries)
	userUsecase := userusecase.NewUserUsecase(userDAO)
	return &UserController{userUsecase: userUsecase}
}