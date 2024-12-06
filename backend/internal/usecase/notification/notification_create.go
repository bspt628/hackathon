package notificationusecase

import (
	"context"
	"hackathon/internal/model"
)

func (rc *NotificationUsecase) CreateNotifications(ctx context.Context, params model.CreateNotificationParams) error {
	
	// リポスト情報をDAOを通じて保存
	return rc.dao.CreateNotifications(ctx, params)
}