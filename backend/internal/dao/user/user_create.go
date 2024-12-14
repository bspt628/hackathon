package userdao

import (
	"context"
	"fmt"
	sqlc "hackathon/db/sqlc/generated"
	"hackathon/internal/auth"
	"hackathon/internal/model"
)

func (dao *UserDAO) CreateUser(ctx context.Context,requestDAO model.UserCreateDAORequest) (*sqlc.User, error) {

	// Firebaseに新規ユーザーを登録
	uid, err := auth.CreateFirebaseUser(requestDAO.Email, requestDAO.Password, requestDAO.Username, requestDAO.DisplayName)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in Firebase: %v", err)
	}

	// UserCreateDAORequest を sqlc.CreateUserParams に変換
	dbParams := model.ToCreateUserParams(requestDAO, uid)

	// SQLクエリを実行して新しいユーザーを作成
	if dao.queries.CreateUser(ctx, dbParams) != nil {
		// DBへの登録に失敗した場合、Firebase登録を削除
		deleteErr := auth.DeleteFirebaseUser(uid)
		if deleteErr != nil {
			// Firebase削除も失敗した場合、ログ出力
			fmt.Printf("failed to rollback Firebase user: %v\n", deleteErr)
		}
		return nil, fmt.Errorf("failed to create user in database: %v", err)
	}

	// 新しく作成されたユーザーの ID で情報を再取得
	user, err := dao.GetUser(ctx, requestDAO.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
