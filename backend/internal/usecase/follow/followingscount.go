package followusecase

import (
	"context"
	"errors"
	"fmt"
	"log"
)

func (fc *FollowUsecase) GetFollowingsCount(ctx context.Context, userID string) (int32, error) {
	if fc == nil || fc.followDAO == nil {
		return 0, errors.New("UserUseCase or UserDAO is nil")
	}
	log.Println("GetFollowingsCount: ", userID)

	// フォロー情報を保存
	count, err := fc.followDAO.GetFollowingsCount(ctx, userID)
	if err != nil {
		return 0, fmt.Errorf("failed to update followinf count: %v", err)
	}

	return count, nil
}