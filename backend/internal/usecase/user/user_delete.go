package usecase

import (
	"context"
)

// DeleteUserByID は、指定されたIDのユーザーを削除する
func (usecase *UserUsecase) DeleteUser(ctx context.Context, id string) error {
	return usecase.dao.DeleteUser(ctx, id)
}