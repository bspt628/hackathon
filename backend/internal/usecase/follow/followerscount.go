package followusecase

import (
	"context"
	"errors"
	"fmt"
)

func (fc *FollowUsecase) GetFollowersCount(ctx context.Context, userID string) (int32, error) {
	if fc == nil || fc.followDAO == nil {
		return 0, errors.New("UserUseCase or UserDAO is nil")
	}

	// フォロー情報を保存
	count, err := fc.followDAO.GetFollowersCount(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to update followers count: %v", err)
	}

	return count, nil
}