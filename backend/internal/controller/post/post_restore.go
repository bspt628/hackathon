package postcontroller

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)



type RestorePostResponse struct {
	Message string `json:"message"`
}

func (pc *PostController) RestorePost(w http.ResponseWriter, r *http.Request) {
	// パスパラメータから post_id を取得
	vars := mux.Vars(r)
	postID := vars["id"]

	// 必須フィールドのチェック
	if postID == "" {
		http.Error(w, "post_id は必須です", http.StatusBadRequest)
		return
	}

	// Usecase層の呼び出し
	err := pc.postUsecase.RestorePost(context.Background(), postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// 成功レスポンス
	response := RestorePostResponse{Message: "投稿が復活しました"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
