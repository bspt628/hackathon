package followcontroller

import(
	"database/sql"
	"hackathon/db/sqlc/generated"
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
	// UserDAOとFollowDAOを作成
	queries := sqlc.New(dbConn)
	userDAO := userdao.NewUserDAO(queries)
	followDAO := followdao.NewFollowDAO(queries)

	userUsecase := userusecase.NewUserUsecase(userDAO)
	followUsecase := followusecase.NewFollowUsecase(followDAO)
	return &FollowController{
		followUsecase: followUsecase,
		userUsecase: userUsecase,
	}
}