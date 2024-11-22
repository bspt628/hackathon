package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

)

func (uc *UserController) UpdateUserBanStatus(w http.ResponseWriter, r *http.Request) {
	var request struct {
		IsBanned bool `json:"is_banned"`
	}
	vars := mux.Vars(r)
	ID := vars["id"]

	if ID == "" {
		http.Error(w, "IDパラメータが指定されていません", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエストの解析に失敗しました: %v", err), http.StatusBadRequest)
		return
	}

	user, err := uc.userUsecase.UpdateUserBanStatus(context.Background(), request.IsBanned, ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザープロフィール更新に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}


}