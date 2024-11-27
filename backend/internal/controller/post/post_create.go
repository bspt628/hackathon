package postcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (pc *PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	var request struct {
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

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエストの解析に失敗しました: %v", err), http.StatusBadRequest)
		return
	}

	post, err := pc.postUsecase.CreatePost(r.Context(), request)
	if err != nil {
		http.Error(w, fmt.Sprintf("投稿の作成に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}