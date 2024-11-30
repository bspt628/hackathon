package userdao

import (
	"context"
	"fmt"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/auth"
)

func (dao *UserDAO) CreateUser(ctx context.Context, arg sqlc.CreateUserParams, password string) (*sqlc.User, error) {
	// arg.Displatnameをstring型に変換
	var displayNameStr string
	if arg.DisplayName.Valid {
		displayNameStr = arg.DisplayName.String
	} else {
		displayNameStr = ""
	}

	// Firebaseに新規ユーザーを登録
	uid, err := auth.CreateFirebaseUser(arg.Email, password, arg.Username, displayNameStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in Firebase: %v", err)
	}
	arg.FirebaseUid = uid
	// fmt.Println(arg)

	// SQLクエリを実行して新しいユーザーを作成
	_, err = dao.queries.CreateUser(ctx, arg)

	if err != nil {
		// DBへの登録に失敗した場合、Firebase登録を削除
		deleteErr := auth.DeleteFirebaseUser(uid)
		if deleteErr != nil {
			// Firebase削除も失敗した場合、ログ出力
			fmt.Printf("failed to rollback Firebase user: %v\n", deleteErr)
		}
		return nil, fmt.Errorf("failed to create user in database: %v", err)
	}

	// 新しく作成されたユーザーの ID で情報を再取得
	user, err := dao.GetUser(ctx, arg.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
