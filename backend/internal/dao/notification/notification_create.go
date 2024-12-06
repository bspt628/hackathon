package notificationdao

import (
	"context"
	"errors"
	"hackathon/internal/model"
)

func (dao *NotificationDAO) CreateNotifications(ctx context.Context, arg model.CreateNotificationParams) error {

	argdao := model.ConvertCreateNotificationsToNotification(arg)

	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("failed to begin transaction")
	}
	defer tx.Rollback()

	// 通知を登録
	err = dao.queries.CreateNotification(ctx, argdao)
	if err != nil {
		return errors.New("failed to create notification")
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return errors.New("failed to commit transaction")
	}

	return nil
}