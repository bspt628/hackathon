package followusecase

import (
	"context"
	"fmt"
	"errors"
	"github.com/google/uuid"
)

func (fc *FollowUsecase) AddFollow(ctx context.Context, firebaseUID, followingID, followerID string) error {
	if fc == nil || fc.followDAO == nil {
        return errors.New("UserUseCase or UserDAO is nil")
    }

	// フォローのビジネスロジックを処理（例えば、自己フォローの禁止など）
	fmt.Println("follow to", followingID, "from",followerID)
	if followerID == followingID {
		return fmt.Errorf("cannot follow yourself")
	}

	// ID生成
	followID := uuid.New().String()

	// フォロー情報をDAOを通じて保存
	return fc.followDAO.AddFollow(ctx, followID, followerID, followingID)
}
