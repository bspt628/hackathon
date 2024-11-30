package followusecase

import (
	"context"
	"fmt"
	"errors"
)

func (fc *FollowUsecase) GetFollowStatus(ctx context.Context, followingID, followerID string) (bool, error) {
	if fc == nil || fc.followDAO == nil {
        return false, errors.New("UserUseCase or UserDAO is nil")
    }

	// フォローのビジネスロジックを処理（例えば、自己フォローの禁止など）
	fmt.Println("[check] follow to", followingID, "from",followerID)
	if followerID == followingID {
		return false, fmt.Errorf("cannot remove following yourself")
	}


	// フォロー情報をDAOを通じて保存
	return fc.followDAO.GetFollowStatus(ctx, followerID, followingID)
}
