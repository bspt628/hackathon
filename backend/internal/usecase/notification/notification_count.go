package notificationusecase

import (
	"context"
	"errors"
)

func (nc *NotificationUsecase) CountUnreadNotifications(ctx context.Context, userID string) (int64, error) {
	// 通知情報をDAOを通じて取得
	notifications, err := nc.dao.CountUnreadNotifications(ctx, userID)
	if err != nil {
		return 0, errors.New("failed to get notifications")
	}
	return notifications, nil
}