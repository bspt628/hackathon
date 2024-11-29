package dao

import (
	"context"
	"fmt"
	sqlc "hackathon/db/sqlc/generated"
	"hackathon/internal/auth"

	"github.com/oklog/ulid"
	"golang.org/x/crypto/bcrypt"
)

func (dao *UserDAO) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (*sqlc.User, error) {
	// IDをulidで自動生成する
	myid := ulid.MustNew(ulid.Now(), nil).String()

	// bcyptでパスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(arg.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// arg.Displatnameをstring型に変換
	var displayNameStr string
	if arg.DisplayName.Valid {
		// displayNameが有効な場合
		displayNameStr = arg.DisplayName.String
	} else {
		// displayNameがNULLの場合
		displayNameStr = ""
	}

	// Firebaseに新規ユーザーを登録
	uid, err := auth.CreateFirebaseUser(arg.Email, arg.PasswordHash, arg.Username, displayNameStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create user in Firebase: %v", err)
	}

	// SQLクエリを実行して新しいユーザーを作成
	_, err = dao.db.CreateUser(ctx, sqlc.CreateUserParams{
		ID:           myid,
		FirebaseUid:  uid,
		Email:        arg.Email,
		PasswordHash: string(hashedPassword),
		Username:     arg.Username,
		DisplayName:  arg.DisplayName,
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