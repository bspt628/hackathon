package userdao

import (
	"context"
	"fmt"
	"hackathon/internal/auth"
)

func (dao *UserDAO) DeleteUser(ctx context.Context, id, firebaseUID string) error {
	// トランザクションの開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// Firebase ユーザーの削除
	deleteErr := auth.DeleteFirebaseUser(firebaseUID)
	if deleteErr != nil {
		// Firebase ユーザーの削除に失敗した場合、ログを記録してトランザクションを中止
		return fmt.Errorf("failed to delete Firebase user: %w", deleteErr)
	}

	// データベースからユーザーを削除
	result, err := dao.queries.DeleteUser(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user from database: %w", err)
	}

	// 削除された行数を確認
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %s", id)
	}

	// トランザクションをコミット
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
