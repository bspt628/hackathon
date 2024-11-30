package userdao

import (
	"context"
	"fmt"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/auth"
	"database/sql"
)

func (dao *UserDAO) CreateUser(ctx context.Context, myid, email, hashedPassword, username, displayname, password string) (*sqlc.User, error) {

	// Firebaseに新規ユーザーを登録
	uid, err := auth.CreateFirebaseUser(email, password, username, displayname)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in Firebase: %v", err)
	}

	// SQLクエリを実行して新しいユーザーを作成
	_, err = dao.queries.CreateUser(ctx, sqlc.CreateUserParams{
		ID: 		  myid,
		FirebaseUid:  uid,
		Email:        email,
		PasswordHash: hashedPassword,
		Username:     username,
		DisplayName:  sql.NullString{String: displayname, Valid: true}, 
	})

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
	user, err := dao.GetUser(ctx, myid)
	if err != nil {
		return nil, err
	}

	return user, nil
}
