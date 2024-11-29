package userusecase

import (
	"context"
	"fmt"
	sqlc "hackathon/db/sqlc/generated"
	"hackathon/domain"
	"hackathon/internal/utils"
	"strings"
)

func (uc *UserUsecase) UpdateUserEmail(ctx context.Context, email string, id string) (*domain.UserUpdateEmailResult, error) {
	// 入力検証
	if !utils.IsValidEmail(email) {
		return nil, fmt.Errorf("有効なメールアドレスを入力してください")
	}

	// todo 現在と同じメールアドレスの場合はエラーを返す

	arg := sqlc.UpdateUserEmailParams{
		Email: email,
		ID:    id,
	}

	// DAO層の関数を呼び出す
	err := uc.dao.UpdateUserEmail(ctx, arg)
	if err != nil {
		// 重複エラーの場合
		if strings.Contains(err.Error(), "is already taken") {
			return nil, fmt.Errorf("このユーザー名は既に使用されています")
		}
		// その他のエラー
		return nil, fmt.Errorf("ユーザー名の更新に失敗しました: %v", err)
	}

	return domain.NewUserUpdateEmailResult(email), nil
}
