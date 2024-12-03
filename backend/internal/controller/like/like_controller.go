package likecontroller

import(
	"database/sql"
	"hackathon/internal/dao/like"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/like"
	"hackathon/internal/usecase/user"
)

type LikeController struct {
	likeUsecase *likeusecase.LikeUsecase
	userUsecase *userusecase.UserUsecase
}

func NewLikeController(dbConn *sql.DB) *LikeController {
	userDAO := userdao.NewUserDAO(dbConn)
	likeDAO := likedao.NewLikeDAO(dbConn)
	userUsecase := userusecase.NewUserUsecase(userDAO)
	likeUsecase := likeusecase.NewLikeUsecase(likeDAO)
	return &LikeController{
		likeUsecase: likeUsecase,
		userUsecase: userUsecase,
	}
}