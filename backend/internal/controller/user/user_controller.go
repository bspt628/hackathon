package usercontroller

import (
	"database/sql"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/user"
)

type UserController struct {
    userUsecase *userusecase.UserUsecase // Usecaseへの依存性注入
}

func NewUserController(dbConn *sql.DB) *UserController {
	userDAO := userdao.NewUserDAO(dbConn)
	userUsecase := userusecase.NewUserUsecase(userDAO)
	return &UserController{userUsecase: userUsecase}
}