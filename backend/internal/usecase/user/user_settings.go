package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"hackathon/db/sqlc/generated"
	"hackathon/domain"
	"hackathon/internal/utils"
	"strings"
	"time"
)

func (uc *UserUsecase) UpdateUserSettings(ctx context.Context, displayName string, birthDate string, language, id string) (*domain.UserSettingsUpdateResult, error) {


	// 更新結果をまとめる
	updatedSettings := map[string]string{}
	if displayName != "" {
		if len(displayName) > 20 {
			return nil, errors.New("表示名は20文字以内で入力してください")
		}
		if strings.TrimSpace(displayName) == "" {
			return nil, errors.New("表示名に空白以外の内容を入力してください")
		}
		updatedSettings["display_name"] = displayName
	}
	// 生年月日のバリデーション
	var birthDateObj sql.NullTime
	if birthDate != "" {
		parsedTime, err := utils.NewDate(birthDate)
		if err != nil {
			return nil, fmt.Errorf("生年月日が無効です: %v", err)
		}
		if parsedTime.After(time.Now()) {
			return nil, errors.New("生年月日は未来の日付を指定できません")
		}
		// UTCに変換してセット
		birthDateObj = sql.NullTime{Time: parsedTime, Valid: true}
		updatedSettings["birth_date"] = parsedTime.Format("2006-01-02")
	}
	if language != "" {
		if len(language) > 100 {
			return nil, errors.New("言語は100文字以内で入力してください")
		}
		updatedSettings["language"] = language
	}

		// UpdateUserProfileParams構造体にデータをセット
		arg := sqlc.UpdateUserSettingsParams{
			DisplayName:    sql.NullString{String: displayName, Valid: displayName != ""},
			BirthDate: 		birthDateObj,
			Language: 		sql.NullString{String: language, Valid: language != ""},
			ID: 			id,
		}

	// ユーザー情報を更新
	err := uc.dao.UpdateUserSettings(ctx, arg)
	if err != nil {
		return nil, err
	}

	return domain.NewUserSettingsUpdateResult(updatedSettings), nil
}
