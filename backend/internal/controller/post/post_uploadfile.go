package postcontroller

import (
	"context"
	"fmt"
	"net/http"
)

func (pc *PostController) UploadFile(w http.ResponseWriter, r *http.Request) {
	// ファイルアップロードの処理をusecase層に依頼
	err := pc.postUsecase.UploadFile(context.Background(), r)
	if err != nil {
		http.Error(w, fmt.Sprintf("ファイルのアップロードに失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "ファイルがアップロードされました"}`))
}
