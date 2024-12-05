package userdao

import (
	"context"
	"database/sql"
	"errors"
	"log"
)

func (dao *UserDAO) GetUserIDByFirebaseUID(ctx context.Context, firebaseUID string) (string, error) {
	log.Println("GetUserIDByFirebaseUID: ", firebaseUID)
	id, err := dao.queries.GetIDfromFirebaseUID(ctx, firebaseUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	return id, nil
}
