package usercontroller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (uc *UserController) GetUserIDByFirebaseUID(w http.ResponseWriter, r *http.Request) {
	// リクエストボディをパース
	vars := mux.Vars(r)
	FirebaseUID := vars["id"]

	userID, err := uc.userUsecase.GetUserIDByFirebaseUID(context.Background(), FirebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get user ID: %v", err), http.StatusInternalServerError)
		return
	}

	// レスポンスを返す
	resp := struct {
		UserID string `json:"user_id"`
	}{
		UserID: userID,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode response: %v", err), http.StatusInternalServerError)
		return
	}
}