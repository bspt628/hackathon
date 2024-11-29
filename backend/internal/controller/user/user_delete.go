package usercontroller

import (
	"context"
	"fmt"
	"net/http"
)

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// FirebaseAuthMiddleware で設定された UserID を取得
	firebaseUID := r.Header.Get("UserID")
	if firebaseUID == "" {
		http.Error(w, "UserID missing in request context", http.StatusUnauthorized)
		return
	}
	id, err := uc.userUsecase.GetUserIDByFirebaseUID(context.Background(), firebaseUID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return 
	}
	fmt.Println("DB id from token",id)

	// Usecase層を通してユーザーを削除
	if err := uc.userUsecase.DeleteUser(context.Background(), id); err != nil {
		http.Error(w, fmt.Sprintf("ユーザー削除に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusNoContent)
}
