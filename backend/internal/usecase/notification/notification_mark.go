package notificationusecase

import (
	"context"
	"errors"
)

func (nc *NotificationUsecase) MarkNotificationsAsRead(ctx context.Context, notificationID string) error {
	// 通知情報をDAOを通じて取得
	err := nc.dao.MarkNotificationsAsRead(ctx, notificationID)
	if err != nil {
		return errors.New("failed to get notifications")
	}
	return nil
}