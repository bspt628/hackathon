package usercontroller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// URLパラメータからユーザーIDを取得
	vars := mux.Vars(r)
	userID := vars["id"]

	if userID == "" {
		http.Error(w, "IDパラメータが指定されていません", http.StatusBadRequest)
		return
	}

	// Usecase層を通してユーザーを削除
	if err := uc.userUsecase.DeleteUser(context.Background(), userID); err != nil {
		http.Error(w, fmt.Sprintf("ユーザー削除に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusNoContent)
}
