package postcontroller

import(
	"database/sql"
	"hackathon/internal/dao/post"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/post"
	"hackathon/internal/usecase/user"
)

type PostController struct {
	postUsecase *postusecase.PostUsecase
	userUsecase *userusecase.UserUsecase
}

func NewPostController(dbConn *sql.DB) *PostController {
	userDAO := userdao.NewUserDAO(dbConn)
	postDAO := postdao.NewPostDAO(dbConn)
	userUsecase := userusecase.NewUserUsecase(userDAO)
	postUsecase := postusecase.NewPostUsecase(postDAO)
	return &PostController{
		postUsecase: postUsecase,
		userUsecase: userUsecase,
	}
}