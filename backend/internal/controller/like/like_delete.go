package likecontroller

import (
	"context"
	"database/sql"
	"net/http"
	"encoding/json"
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

	err = lc.likeUsecase.DeleteLike(context.Background(), request.PostID, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Like not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete like", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Deleted like successfully"}`))
}