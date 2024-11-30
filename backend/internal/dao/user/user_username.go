package userdao

import (
	"context"
	"fmt"
	"hackathon/db/sqlc/generated"

	"github.com/go-sql-driver/mysql"
)

// ユーザーのメールアドレスを更新
func (dao *UserDAO) UpdateUserName(ctx context.Context, params sqlc.UpdateUserNameParams) error {
	err := dao.queries.UpdateUserName(ctx, params)

	if err != nil {
		// MySQLエラーを検出
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 { // ER_DUP_ENTRY
				return fmt.Errorf("email '%s' is already taken", params.Username)
			}
		}
		// その他のエラー
		return fmt.Errorf("failed to update username: %v", err)
	}
	return nil
}