package likecontroller

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func (lc *LikeController) DeleteLike(w http.ResponseWriter, r *http.Request) {
	firebaseID := r.Header.Get("UserID")
	if firebaseID == "" {
		http.Error(w, "UserID missing in request context", http.StatusUnauthorized)
		return
	}

	userID, err := lc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseID)
	if err != nil {
		http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	postID := vars["id"]

	// 必須パラメータをチェック
	if postID == "" {
		http.Error(w, "PostID is required", http.StatusBadRequest)
		return 
	}

	err = lc.likeUsecase.DeleteLike(context.Background(), postID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Like not found", http.StatusNotFound)
			return
		}
		// errを返す
		http.Error(w, fmt.Sprintf("いいねの削除に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Deleted like successfully"}`))
}