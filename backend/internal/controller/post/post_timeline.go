package postcontroller

import (
	"net/http"
	"strconv"
	"context"
	"fmt"
	"encoding/json"
)

func (pc *PostController) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // デフォルト値
	}

    posts, err := pc.postUsecase.GetAllPosts(context.Background(), int32(limit))
    if err != nil {
		http.Error(w, fmt.Sprintf("投稿の取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}


    w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
	
}

func (pc *PostController) GetFollowedUsersPosts(w http.ResponseWriter, r *http.Request) {
    firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		http.Error(w, "UserID missing in request context", http.StatusUnauthorized)
		return
	}

	userID, err := pc.userUsecase.GetUserIDFromFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("フォローするユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	limitStr := r.URL.Query().Get("limit")
    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit <= 0 {
        limit = 10 // デフォルト値
    }

    posts, err := pc.postUsecase.GetFollowedUsersPosts(context.Background(), userID, int32(limit))
    if err != nil {
		http.Error(w, fmt.Sprintf("フォロー中のユーザーのポストの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

    w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}