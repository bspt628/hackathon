package postcontroller

import(
	"database/sql"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/dao/post"
	"hackathon/internal/usecase/post"
)

type PostController struct {
	postUsecase *postusecase.PostUsecase
}

func NewPostController(dbConn *sql.DB) *PostController {
	queries := sqlc.New(dbConn)
	postDAO := postdao.NewPostDAO(queries)
	postUsecase := postusecase.NewPostUsecase(postDAO)
	return &PostController{postUsecase: postUsecase}
}