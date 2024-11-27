package usecase

import (
	"context"
	"fmt"
	"strings"
	"hackathon/db/sqlc/generated"
	
	"hackathon/domain"
)


func (uc *UserUsecase) UpdateUserName(ctx context.Context, username string, id string)  (*domain.UserUpdateUsernameResult, error) {
	// 入力検証
	if username == "" {
		return nil, fmt.Errorf("ユーザー名を入力してください")
	}

	// todo 現在と同じメールアドレスの場合はエラーを返す
		

	arg := sqlc.UpdateUserNameParams{
		Username: username,
		ID:			  id,
	}

	// DAO層の関数を呼び出す
	err := uc.dao.UpdateUserName(ctx, arg)
	if err != nil {
		// 重複エラーの場合
		if strings.Contains(err.Error(), "is already taken") {
			return nil, fmt.Errorf("このユーザー名は既に使用されています")
		}
		// その他のエラー
		return nil, fmt.Errorf("ユーザー名の更新に失敗しました: %v", err)
	}

	return domain.NewUserUpdateUsernameResult(username), nil
}