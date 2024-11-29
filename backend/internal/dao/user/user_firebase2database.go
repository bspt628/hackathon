package dao

import (
	"context"
	"database/sql"
	"errors"
)

// GetUserIDByFirebaseUID は Firebase UID に対応するユーザー ID を取得します
func (dao *UserDAO) GetUserIDByFirebaseUID(ctx context.Context, firebaseUID string) (string, error) {
	id, err := dao.db.GetIdfromFirebaseUID(ctx, firebaseUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	return id, nil
}
