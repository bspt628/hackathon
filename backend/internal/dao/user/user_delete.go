package userdao

import (
	"context"
	"fmt"
)

func (dao *UserDAO) DeleteUser(ctx context.Context, id string) error {
	result, err := dao.db.DeleteUser(ctx, id)
	if err != nil {
		return err // SQL 実行エラーをそのまま返す
	}

	rowsAffected, err := result.RowsAffected()
	fmt.Println(result)
	fmt.Println(rowsAffected)
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %s", id) // 削除対象がない場合のエラー
	}

	return nil
}
