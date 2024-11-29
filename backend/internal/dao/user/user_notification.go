package userdao

import (
	"context"
	sqlc "hackathon/db/sqlc/generated"
)

func (dao *UserDAO) UpdateUserNotifications(ctx context.Context, params sqlc.UpdateUserNotificationsParams) error {
	return dao.db.UpdateUserNotifications(ctx, params)
}
