package likecontroller

import (
	"context"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

func (lc *LikeController) GetPostLikesCount(w http.ResponseWriter, r *http.Request) {
	// パスパラメータから投稿IDを取得
	vars := mux.Vars(r)
	postID := vars["id"]

	count, err := lc.likeUsecase.GetPostLikesCount(context.Background(), postID)
	if err != nil{
		http.Error(w, fmt.Sprintf("いいねの数の取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(count); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}