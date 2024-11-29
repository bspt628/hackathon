package userdao

import (
	"context"
	"fmt"
)

func (dao *UserDAO) GetEmailByUsername(ctx context.Context, username string) (string, error) {
	email, err := dao.db.GetEmailFromUsername(ctx, username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" { // SQLCでは明示的なエラーは生成しない
			return "", fmt.Errorf("no user found with username: %s", username)
		}
		return "", fmt.Errorf("failed to fetch email for username %s: %v", username, err)
	}
	return email, nil
}
