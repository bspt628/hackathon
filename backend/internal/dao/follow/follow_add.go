package followdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"github.com/go-sql-driver/mysql"
	"errors"
)

func (dao *FollowDAO) AddFollow(ctx context.Context, followID, followerID, followingID string) error {
	// トランザクションを開始
	tx, err := dao.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// フォロー情報を保存
	err = dao.queries.AddFollow(ctx, sqlc.AddFollowParams{
		ID: followID, 
		FollowerID: sql.NullString{String: followerID, Valid: true},
		FollowingID: sql.NullString{String: followingID, Valid: true},
	})
	if err != nil {
		if errmysql, ok := err.(*mysql.MySQLError); ok {
			if errmysql.Number == 1062 {
				return errors.New("follow already exists")
			}
		}else {
			return errors.New("failed to add follow")
		}
	}
	// フォロワー数をインクリメント
	_, err = dao.queries.IncrementFollowersCount(ctx, followingID)
	if err != nil {
		return errors.New("failed to IncrementFollowersCount")
	}
	// フォロー数をインクリメント		
	_, err = dao.queries.IncrementFollowingsCount(ctx, followerID)
	if err != nil {
		return errors.New("failed to IncrementFollowingsCount")
	}

	// トランザクションをコミット
	err = tx.Commit()
	if err != nil {
		return errors.New("failed to commit transaction")
	}
	
	return nil
}