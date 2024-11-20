package usecase

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
)

func (uc *UserUsecase) UpdateUserProfile(ctx context.Context, profile_image_url, bio, location, website, id string) (*UserProfileUpdateResult, error) {
	// UpdateUserProfileParams構造体にデータをセット
	arg := sqlc.UpdateUserProfileParams{
		ProfileImageUrl: sql.NullString{String: profile_image_url, Valid: profile_image_url != ""},
		Bio:             sql.NullString{String: bio, Valid: bio != ""},
		Location:        sql.NullString{String: location, Valid: location != ""},
		Website:         sql.NullString{String: website, Valid: website != ""},
		ID:			  id,
	}

	// ユーザー情報を更新
	err := uc.dao.UpdateUserProfile(ctx, arg)
	if err != nil {
		return nil, err
	}

	// 更新結果をまとめる
	updatedFields := map[string]string{}
	if arg.ProfileImageUrl.Valid {
		updatedFields["profile_image_url"] = arg.ProfileImageUrl.String
	}
	if arg.Bio.Valid {
		updatedFields["bio"] = arg.Bio.String
	}
	if arg.Location.Valid {
		updatedFields["location"] = arg.Location.String
	}
	if arg.Website.Valid {
		updatedFields["website"] = arg.Website.String
	}

	return NewUserProfileUpdateResult(updatedFields), nil
}
