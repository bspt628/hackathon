package domain

type CreatePostRequest struct {
	UserID         string   `json:"user_id"`
	Content        string   `json:"content"`
	MediaURLs      []string `json:"media_urls"`
	Visibility     string   `json:"visibility"`
	Language       string   `json:"language"`
	Location       string   `json:"location"`
	Device         string   `json:"device"`
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
	Language   string   `json:"language"`
	Location   string   `json:"location"`
	Device     string   `json:"device"`
}