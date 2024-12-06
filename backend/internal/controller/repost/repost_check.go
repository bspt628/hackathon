package repostcontroller

import (
	"context"
	"hackathon/internal/model"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

)

func (rc RepostController) GetRepostStatus(w http.ResponseWriter, r *http.Request) {
	firebaseUID := r.Header.Get("UserID")

	userID, err := rc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, "ユーザーIDの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	postID := vars["id"]

	params := model.GetRepostStatusParams{
		UserID: userID,
		OriginalPostID: postID,
	}

	result, err := rc.repostUsecase.GetRepostStatus(context.Background(), params)
	if err != nil {
		http.Error(w, "リポストのステータスの取得に失敗しました", http.StatusInternalServerError)
		return 
	}

	response := map[string]bool{"repost_status": result}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}
