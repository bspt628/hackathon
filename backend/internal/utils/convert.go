package utils

import (
	"database/sql"
	"time"
)

// 補助関数: sql.Null型をGoの基本型に変換
func ConvertNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func ConvertNullTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{} // デフォルト値
}

func ConvertNullBool(nb sql.NullBool) bool {
	if nb.Valid {
		return nb.Bool
	}
	return false
}

func ConvertNullInt32(ni sql.NullInt32) int32 {
	if ni.Valid {
		return ni.Int32
	}
	return 0
}