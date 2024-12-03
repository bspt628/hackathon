package postusecase

import (
	"hackathon/internal/dao/post"
)

type PostUsecase struct {
	dao *postdao.PostDAO
}

func NewPostUsecase(dao *postdao.PostDAO) *PostUsecase {
	return &PostUsecase{dao: dao}
}