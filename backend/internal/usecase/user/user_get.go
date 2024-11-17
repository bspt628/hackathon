package usecase

import (
	"context"
	"hackathon/db/sqlc/generated"
)

func (usecase *UserUsecase) GetUser(ctx context.Context, id string) (*sqlc.User, error) {
	user, err := usecase.dao.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
