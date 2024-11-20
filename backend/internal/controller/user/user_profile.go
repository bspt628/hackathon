package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

)

func (uc *UserController) UpdateUserProfile(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ProfileImageUrl string `json:"profile_image_url"`
		Bio             string `json:"bio"`
		Location        string `json:"location"`
		Website         string `json:"website"`
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

	if request.ProfileImageUrl == "" && request.Bio == "" && request.Location == "" && request.Website == "" {
		http.Error(w, "何か変更をしてください", http.StatusBadRequest)
		return
	}

	user, err := uc.userUsecase.UpdateUserProfile(context.Background(), request.ProfileImageUrl, request.Bio, request.Location, request.Website, ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザープロフィール更新に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}


}