package domain

import (
	"encoding/json"
)

type UserProfileUpdateResult struct {
	UpdatedFields map[string]string `json:"updated_fields"`
}

func NewUserProfileUpdateResult(updatedFields map[string]string) *UserProfileUpdateResult {
	return &UserProfileUpdateResult{
		UpdatedFields: updatedFields,
	}
}

type UserSettingsUpdateResult struct {
	UpdatedSettings map[string]string `json:"updated_settings"`
}

func NewUserSettingsUpdateResult(updatedSettings map[string]string) *UserSettingsUpdateResult {
	return &UserSettingsUpdateResult{
		UpdatedSettings: updatedSettings,
	}
}

type UserNotificationsUpdateResult struct {
	UpdatedNotifications json.RawMessage `json:"updated_notifications"`
}

func NewUserNotificationsUpdateResult(updatedNotifications json.RawMessage) *UserNotificationsUpdateResult {
	return &UserNotificationsUpdateResult{
		UpdatedNotifications: updatedNotifications,
	}
}

type UserPrivacyUpdateResult struct {
	IsPrivate bool `json:"is_private"`
}

func NewUserPrivacyUpdateResult(isPrivate bool) *UserPrivacyUpdateResult {
	return &UserPrivacyUpdateResult{
		IsPrivate: isPrivate,
	}
}

type UserBanStatusUpdateResult struct {
	IsBanned bool `json:"is_banned"`
}

func NewUserBanStatusUpdateResult(isBanned bool) *UserBanStatusUpdateResult {
	return &UserBanStatusUpdateResult{
		IsBanned: isBanned,
	}
}

type NotificationSettings struct {
	Enabled   bool   `json:"enabled"`
	Frequency string `json:"frequency"`
}