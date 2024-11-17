package usecase

import (
	"context"
	"hackathon/db/sqlc/generated"
)

func (usecase *UserUsecase) GetUserByID(ctx context.Context, id string) (*db.User, error) {
	user, err := usecase.dao.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
