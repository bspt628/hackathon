package utils

import (
	"database/sql"
	"time"
	"fmt"
)

type Date struct {
	sql.NullTime
}

// Date型のコンストラクタメソッド
func NewDate(birthDate string) (time.Time, error) {
	// birthDateが空文字でない場合のみ、time.TimeをセットしてValidをtrueにする
	if birthDate == "" {
		return time.Time{}, fmt.Errorf("誕生日が設定されていません") // 空文字の場合はNullTimeを無効にする
	}
	// 日付文字列をtime.Timeに変換
	return  time.Parse("2006-01-02", birthDate)
}