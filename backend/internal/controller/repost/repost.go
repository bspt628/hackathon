package repostcontroller

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"
	"hackathon/internal/model"
	"github.com/google/uuid"
)



func (rc *RepostController) CreateRepost(w http.ResponseWriter, r *http.Request) {
	// FirebaseAuthMiddleware で設定された UserID を取得
	firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		http.Error(w, "UserID missing in request context", http.StatusUnauthorized)
		return
	}

	userID, err := rc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("リポストするユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}
	
	// HTTPリクエストのjsonをdecodeしてパラメータ「repostID」を取得
	var request struct {
		PostID string `json:"post_id"`
		IsQuoteRepost bool `json:"is_quote_repost"`
		AdditionalComment string `json:"additional_comment"`
	}

	// リクエストのJSONデータを構造体にバインド
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエストの解析に失敗しました: %v", err), http.StatusBadRequest)
		return
	}

	// 必須パラメータをチェック
	if request.PostID == "" {
		http.Error(w, "PostID is required", http.StatusBadRequest)
		return 
	}

	params := model.CreateRepostParams{
		ID:                uuid.New().String(),
		UserID:            userID,
		OriginalPostID:    request.PostID,
		IsQuoteRepost:     request.IsQuoteRepost,
		AdditionalComment: request.AdditionalComment,
	}

	// リポストを実行
	err = rc.repostUsecase.CreateRepost(context.Background(), params)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to repost: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Reposted successfully"}`))
}