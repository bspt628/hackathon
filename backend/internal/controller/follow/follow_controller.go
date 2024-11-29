package followcontroller

import(
	"database/sql"
	"hackathon/internal/dao/follow"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/follow"
	"hackathon/internal/usecase/user"
)

type FollowController struct {
	followUsecase *followusecase.FollowUsecase
	userUsecase *userusecase.UserUsecase
}

func NewFollowController(dbConn *sql.DB) *FollowController {
	userDAO := userdao.NewUserDAO(dbConn)
	followDAO := followdao.NewFollowDAO(dbConn)

	userUsecase := userusecase.NewUserUsecase(userDAO)
	followUsecase := followusecase.NewFollowUsecase(followDAO)
	return &FollowController{
		followUsecase: followUsecase,
		userUsecase: userUsecase,
	}
}