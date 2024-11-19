package usecase

import (
	"context"
	"hackathon/db/sqlc/generated"
)

func (uc *UserUsecase) GetUser(ctx context.Context, id string) (*sqlc.User, error) {
	user, err := uc.dao.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
