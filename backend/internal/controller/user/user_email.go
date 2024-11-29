package usercontroller

import (
	"context"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

func (uc *UserController) UpdateUserEmail(w http.ResponseWriter, r *http.Request){
	var request struct {
		Email string `json:"email"`
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

	email, err := uc.userUsecase.UpdateUserEmail(context.Background(), request.Email, ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザープロフィール更新に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(email); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}