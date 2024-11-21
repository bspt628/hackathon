package usecase

import (
	"context"
	"hackathon/db/sqlc/generated"
	"hackathon/domain"
	"encoding/json"
	"fmt"

)




func (uc *UserUsecase) UpdateUserNotifications(ctx context.Context, notificationSettings domain.NotificationSettings, id string) (*domain.UserNotificationsUpdateResult, error) {
	// notificationSettings を JSON にエンコード
	rawSettings, err := json.Marshal(notificationSettings)
	if err != nil {
		return nil, fmt.Errorf("通知設定のエンコードに失敗しました: %v", err)
	}

	arg := sqlc.UpdateUserNotificationsParams{
		NotificationSettings: rawSettings, // JSON形式のバイト列
		ID:                   id,
	}

	// DAO層で更新処理
	if err := uc.dao.UpdateUserNotifications(ctx, arg); err != nil {
		return nil, fmt.Errorf("通知設定の更新に失敗しました: %v", err)
	}

	return domain.NewUserNotificationsUpdateResult(arg.NotificationSettings), nil
}
