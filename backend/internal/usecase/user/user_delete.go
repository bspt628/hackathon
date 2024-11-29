package userusecase

import (
	"context"
)

// DeleteUserByID は、指定されたIDのユーザーを削除する
func (uc *UserUsecase) DeleteUser(ctx context.Context, id string) error {
	return uc.dao.DeleteUser(ctx, id)
}
