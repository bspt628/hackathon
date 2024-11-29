package userdao

import (
	"context"
	"fmt"
	sqlc "hackathon/db/sqlc/generated"

	"github.com/go-sql-driver/mysql"
)

// ユーザーのメールアドレスを更新
func (dao *UserDAO) UpdateUserEmail(ctx context.Context, params sqlc.UpdateUserEmailParams) error {
	err := dao.db.UpdateUserEmail(ctx, params)

	if err != nil {
		// MySQLエラーを検出
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 { // ER_DUP_ENTRY
				return fmt.Errorf("email '%s' is already taken", params.Email)
			}
		}
		// その他のエラー
		return fmt.Errorf("failed to update email: %v", err)
	}
	return nil
}
