package userusecase

import (
	"context"
	"errors"
)

// GetUserIDByFirebaseUID は Firebase UID からユーザー ID を取得します
func (uc *UserUsecase) GetUserIDByFirebaseUID(ctx context.Context, firebaseUID string) (string, error) {
	userID, err := uc.dao.GetUserIDByFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		return "", errors.New("failed to fetch user ID: " + err.Error())
	}
	return userID, nil
}