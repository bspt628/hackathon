package usecase

import (
	"context"
	"database/sql"
	"hackathon/db/sqlc/generated"
	"strings"
	"errors"
	"hackathon/internal/utils"
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

	// 更新結果をまとめる
	updatedFields := map[string]string{}
	if arg.ProfileImageUrl.Valid {
		updatedFields["profile_image_url"] = arg.ProfileImageUrl.String
		if !utils.IsValidURL(profile_image_url) {
			return nil, errors.New("プロフィール画像のURLが無効です")
		}
		if !strings.HasPrefix(profile_image_url, "http://") && !strings.HasPrefix(profile_image_url, "https://") {
			return nil, errors.New("プロフィール画像のURLはHTTPまたはHTTPSである必要があります")
		}
	}
	if arg.Bio.Valid {
		updatedFields["bio"] = arg.Bio.String
		if len(bio) > 500 {
			return nil, errors.New("自己紹介は500文字以内で入力してください")
		}
		if strings.TrimSpace(bio) == "" {
			return nil, errors.New("自己紹介に空白以外の内容を入力してください")
		}
	}
	if arg.Location.Valid {
		updatedFields["location"] = arg.Location.String
		if len(location) > 100 {
			return nil, errors.New("居住地は100文字以内で入力してください")
		}
	}
	if arg.Website.Valid {
		updatedFields["website"] = arg.Website.String
		if len(website) > 255 {
			return nil, errors.New("ウェブサイトのURLは255文字以内で入力してください")
		}
		if !utils.IsValidURL(website) {
			return nil, errors.New("ウェブサイトのURLが無効です")
		}
		if !strings.HasPrefix(website, "http://") && !strings.HasPrefix(website, "https://") {
			return nil, errors.New("ウェブサイトのURLはHTTPまたはHTTPSである必要があります")
		}
	}

	// ユーザー情報を更新
	err := uc.dao.UpdateUserProfile(ctx, arg)
	if err != nil {
		return nil, err
	}

	return NewUserProfileUpdateResult(updatedFields), nil
}
