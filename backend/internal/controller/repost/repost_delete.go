package repostcontroller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"hackathon/internal/model"
)

func (rc *RepostController) DeleteRepost(w http.ResponseWriter, r *http.Request) {
	firebaseID := r.Header.Get("UserID")
	if firebaseID == "" {
		http.Error(w, "UserID missing in request context", http.StatusUnauthorized)
		return
	}

	userID, err := rc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseID)
	if err != nil {
		http.Error(w, "Failed to get user ID", http.StatusInternalServerError)
		return
	}

	var request struct {
		PostID string `json:"post_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	if request.PostID == "" {
		http.Error(w, "PostID is required", http.StatusBadRequest)
		return
	}

	params := model.DeleteRepostParams{
		UserID:            userID,
		OriginalPostID:    request.PostID,
	}

	err = rc.repostUsecase.DeleteRepost(context.Background(), params)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Repost not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("リポストの削除に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Deleted repost successfully"}`))
}