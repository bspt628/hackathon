package notificationdao

import (
	"context"
	"fmt"
)

func (dao *NotificationDAO) MarkNotificationsAsRead(ctx context.Context, notificationID string) error {
	result, err := dao.queries.MarkNotificationsAsRead(ctx, notificationID)
	if err != nil {
		return fmt.Errorf("failed to mark notifications as read: %w", err)
	}
	// 既読になった通知の数を確認
	count, err := result.RowsAffected() 
	if err != nil {
		return fmt.Errorf("failed to get the number of notifications that have been read: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("no notifications have been read: %w", err)
	}

	return nil
}