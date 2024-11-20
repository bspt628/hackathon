package dao

import (
	"context"
	"fmt"
)

// UserDAO構造体とNewUserDAOは既に定義済み

// GetEmailByUsername メソッド: ユーザー名からメールアドレスを取得
func (dao *UserDAO) GetEmailByUsername(ctx context.Context, username string) (string, error) {
	// sqlcで生成された関数を呼び出し
	email, err := dao.db.GetEmailFromUsername(ctx, username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" { // SQLCでは明示的なエラーは生成しない
			return "", fmt.Errorf("no user found with username: %s", username)
		}
		return "", fmt.Errorf("failed to fetch email for username %s: %v", username, err)
	}
	return email, nil
}
