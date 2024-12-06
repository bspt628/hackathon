package userusecase

import (
	"context"
	"errors"
	"log"
)

func (uc *UserUsecase) GetUserIDByFirebaseUID(ctx context.Context, firebaseUID string) (string, error) {
	log.Println("GetUserIDByFirebaseUID: ", firebaseUID)
	userID, err := uc.dao.GetUserIDByFirebaseUID(ctx, firebaseUID)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	return userID, nil
}