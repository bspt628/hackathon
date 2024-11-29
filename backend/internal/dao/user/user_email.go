package userdao

import (
	"context"
	"fmt"
	sqlc "hackathon/db/sqlc/generated"

	"github.com/go-sql-driver/mysql"
)

func (dao *UserDAO) UpdateUserEmail(ctx context.Context, params sqlc.UpdateUserEmailParams) error {
	err := dao.queries.UpdateUserEmail(ctx, params)

	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 { // ER_DUP_ENTRY
				return fmt.Errorf("email '%s' is already taken", params.Email)
			}
		}
		return fmt.Errorf("failed to update email: %v", err)
	}
	return nil
}
