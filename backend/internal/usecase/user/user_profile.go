package userusecase

import (
	"context"
	"database/sql"
	"errors"
	"hackathon/db/sqlc/generated"
	"hackathon/domain"
	"net/url"
	"strings"
)

func (uc *UserUsecase) UpdateUserProfile(ctx context.Context, profile_image_url, bio, location, website, id string) (*domain.UserProfileUpdateResult, error) {
	// UpdateUserProfileParams構造体にデータをセット
	arg := sqlc.UpdateUserProfileParams{
		ProfileImageUrl: sql.NullString{String: profile_image_url, Valid: profile_image_url != ""},
		Bio:             sql.NullString{String: bio, Valid: bio != ""},
		Location:        sql.NullString{String: location, Valid: location != ""},
		Website:         sql.NullString{String: website, Valid: website != ""},
		ID:              id,
	}

	// 更新結果をまとめる
	updatedFields := map[string]string{}
	if arg.ProfileImageUrl.Valid {
		updatedFields["profile_image_url"] = arg.ProfileImageUrl.String
		if !IsValidURL(profile_image_url) {
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
		if !IsValidURL(website) {
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

	return domain.NewUserProfileUpdateResult(updatedFields), nil
}

func IsValidURL(rawURL string) bool {
	// 文字列が空の場合は無効
	if rawURL == "" {
		return false
	}

	// URLを解析
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return false
	}

	// スキーム（httpまたはhttps）とホスト名が存在するかを確認
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return false
	}
	if parsedURL.Host == "" {
		return false
	}

	return true
}
