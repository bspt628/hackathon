package utils

import (
	"net/url"
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
