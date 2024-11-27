package utils

import (
	"net/url"
	"regexp"
)

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

// メールアドレスの形式を検証する関数
func IsValidEmail(email string) bool {
	// メールアドレス用の正規表現
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}