package followusecase

import (
	"context"
	"fmt"
	"errors"
)

func (fc *FollowUsecase) RemoveFollow(ctx context.Context, followingID, followerID string) error {
	if fc == nil || fc.followDAO == nil {
        return errors.New("UserUseCase or UserDAO is nil")
    }

	// フォローのビジネスロジックを処理（例えば、自己フォローの禁止など）
	fmt.Println("[remove] follow to", followingID, "from",followerID)
	if followerID == followingID {
		return fmt.Errorf("cannot remove following yourself")
	}


	// フォロー情報をDAOを通じて保存
	return fc.followDAO.RemoveFollow(ctx, followerID, followingID)
}
