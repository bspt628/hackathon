package model

import (
	"hackathon/db/sqlc/generated"
	"database/sql"
)

func ConvertCreateRepostParamsToRepost(arg CreateRepostParams) sqlc.CreateRepostParams {
	return sqlc.CreateRepostParams{
		ID: arg.ID,
		UserID: sql.NullString{String: arg.UserID, Valid: true},
		OriginalPostID: sql.NullString{String: arg.OriginalPostID, Valid: true},
		IsQuoteRepost: sql.NullBool{Bool: arg.IsQuoteRepost, Valid: true},
		AdditionalComment: sql.NullString{String: arg.AdditionalComment, Valid: true},
	}
}

func ConvertDeleteRepostParamsToRepost(arg DeleteRepostParams) sqlc.DeleteRepostParams {
	return sqlc.DeleteRepostParams{
		UserID: sql.NullString{String: arg.UserID,Valid: true},
		OriginalPostID: sql.NullString{String: arg.OriginalPostID, Valid: true},
	}
}

func ConvertGetRepostStatusParamsToRepost(arg GetRepostStatusParams) sqlc.GetRepostStatusParams {
	return sqlc.GetRepostStatusParams{
		UserID: sql.NullString{String: arg.UserID, Valid: true},
		OriginalPostID: sql.NullString{String: arg.OriginalPostID, Valid: true},
	}
}

func ConvertCreateNotificationsToNotification(arg CreateNotificationParams) sqlc.CreateNotificationParams {
	return sqlc.CreateNotificationParams{
		ID: arg.ID,
		UserID: sql.NullString{String: arg.UserID, Valid: true},
		Type: sql.NullString{String: arg.Type, Valid: true},
		Message: sql.NullString{String: arg.Message, Valid: true},
	}
}

