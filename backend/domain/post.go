package domain

import (
	"encoding/json"
	"time"
)

type CreatePostRequest struct {
	UserID         string   `json:"user_id"`
	Content        string   `json:"content"`
	MediaURLs      []string `json:"media_urls"`
	Visibility     string   `json:"visibility"`
	OriginalPostID *string  `json:"original_post_id"`
	ReplyToID      *string  `json:"reply_to_id"`
	RootPostID     *string  `json:"root_post_id"`
	IsRepost       bool     `json:"is_repost"`
	IsReply        bool     `json:"is_reply"`
}

type Post struct {
	ID         string   `json:"id"`
	UserID     string   `json:"user_id"`
	Content    string   `json:"content"`
	MediaURLs  []string `json:"media_urls"`
	Visibility string   `json:"visibility"`
}

type PostAll struct {
	ID             string          `json:"id"`
	UserID         string          `json:"user_id"`
	Content        string          `json:"content"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	IsRepost       bool            `json:"is_repost"`
	OriginalPostID string          `json:"original_post_id"`
	ReplyToID      string          `json:"reply_to_id"`
	RootPostID     string          `json:"root_post_id"`
	IsReply        bool            `json:"is_reply"`
	MediaUrls      json.RawMessage `json:"media_urls"`
	LikesCount     int32           `json:"likes_count"`
	RepostsCount   int32           `json:"reposts_count"`
	RepliesCount   int32           `json:"replies_count"`
	ViewsCount     int32           `json:"views_count"`
	Visibility     string          `json:"visibility"`
	IsPinned       bool            `json:"is_pinned"`
	IsDeleted      bool            `json:"is_deleted"`
	Username       string          `json:"username"`
	DisplayName    string          `json:"display_name"`
}
