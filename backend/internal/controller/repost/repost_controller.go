package repostcontroller

import (
	"database/sql"
	"hackathon/internal/dao/repost"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/repost"
	"hackathon/internal/usecase/user"
)

type RepostController struct {
	repostUsecase *repostusecase.RepostUsecase
	userUsecase   *userusecase.UserUsecase
}

func NewRepostController(dbConn *sql.DB) *RepostController {
	userDAO := userdao.NewUserDAO(dbConn)
	repostDAO := repostdao.NewRepostDAO(dbConn)
	userUsecase := userusecase.NewUserUsecase(userDAO)
	repostUsecase := repostusecase.NewRepostUsecase(repostDAO)
	return &RepostController{
		repostUsecase: repostUsecase,
		userUsecase:   userUsecase,
	}
}


