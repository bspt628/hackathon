package userdao

import (
	"context"
)

func (dao *UserDAO) DeleteUser(ctx context.Context, id string) error {
	return dao.db.DeleteUser(ctx, id)
}
