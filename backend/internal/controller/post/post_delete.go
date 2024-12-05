package postcontroller

import (
	"context"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func (pc *PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
	// URLパスからパラメータ「username」を取得
	vars := mux.Vars(r)
	postId := vars["id"]

	// 必須パラメータをチェック
	if postId == "" {
		http.Error(w, "followingid is required", http.StatusBadRequest)
		return 
	}


	// フォローを実行
	err := pc.postUsecase.DeletePost(context.Background(), postId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to remove post: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Follow removed successfully"}`))


}