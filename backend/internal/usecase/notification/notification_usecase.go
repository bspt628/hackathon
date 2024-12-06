package notificationusecase

import (
	"hackathon/internal/dao/notification"
)

type NotificationUsecase struct {
	dao *notificationdao.NotificationDAO
}

func NewNotificationUsecase(dao *notificationdao.NotificationDAO) *NotificationUsecase {
	return &NotificationUsecase{dao: dao}
}