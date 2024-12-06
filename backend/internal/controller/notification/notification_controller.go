package notificationcontroller

import (
	"database/sql"
	"hackathon/internal/dao/notification"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/notification"
	"hackathon/internal/usecase/user"
)

type NotificationController struct{
	notificationUsecase *notificationusecase.NotificationUsecase
	userUsecase *userusecase.UserUsecase
}

func NewNotificationController(dbConn *sql.DB) *NotificationController {
	userDAO := userdao.NewUserDAO(dbConn)
	notificationDAO := notificationdao.NewNotificationDAO(dbConn)
	userUsecase := userusecase.NewUserUsecase(userDAO)
	notificationUsecase := notificationusecase.NewNotificationUsecase(notificationDAO)
	return &NotificationController{
		notificationUsecase: notificationUsecase,
		userUsecase: userUsecase,
	}
}