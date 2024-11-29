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
}

func NewFollowController(dbConn *sql.DB) *FollowController {
	// UserDAOとFollowDAOを作成
	queries := sqlc.New(dbConn)
	userDAO := dao.NewUserDAO(queries)
	followDAO := followdao.NewFollowDAO(queries)

	// UserUsecaseを初期化
	userUsecase := usecase.NewUserUsecase(userDAO)

	// FollowUsecaseを初期化
	followUsecase := followusecase.NewFollowUsecase(followDAO, userUsecase)
	return &FollowController{followUsecase: followUsecase}
}