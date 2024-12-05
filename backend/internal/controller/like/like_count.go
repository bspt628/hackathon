package likecontroller

import (
	"context"
	"net/http"
	"encoding/json"
	"fmt"
)

func (lc *LikeController) GetPostLikesCount(w http.ResponseWriter, r *http.Request) {
	// HTTPリクエストのjsonをdecodeしてパラメータ「likeID」を取得
	var request struct {
		PostID string `json:"post_id"`
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

	count, err := lc.likeUsecase.GetPostLikesCount(context.Background(), request.PostID)
	if err != nil{
		http.Error(w, fmt.Sprintf("いいねの数の取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(count); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}