package notificationdao

import (
	"context"
	"database/sql"
)

func (dao *NotificationDAO) CountUnreadNotifications(ctx context.Context, userID string) (int64, error) {
	userIDDAO := sql.NullString{String: userID, Valid: true}
	count, err := dao.queries.CountUnreadNotifications(ctx, userIDDAO)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}

func (dao *NotificationDAO) CountAllNotifications(ctx context.Context, userID string) (int64, error) {
	userIDDAO := sql.NullString{String: userID, Valid: true}
	count, err := dao.queries.CountAllNotifications(ctx, userIDDAO)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}