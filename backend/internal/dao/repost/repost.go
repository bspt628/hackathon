package repostdao

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/model"
)

func (dao *RepostDAO) CreateRepost(ctx context.Context, arg model.CreateRepostParams) error {
	argdao := sqlc.CreateRepostParams{
		ID: arg.ID,
		UserID: sql.NullString{String: arg.UserID, Valid: true},
		OriginalPostID: sql.NullString{String: arg.OriginalPostID, Valid: true},
		IsQuoteRepost: sql.NullBool{Bool: arg.IsQuoteRepost, Valid: true},
		AdditionalComment: sql.NullString{String: arg.AdditionalComment, Valid: true},
	}

	return dao.queries.CreateRepost(ctx, argdao)
}

