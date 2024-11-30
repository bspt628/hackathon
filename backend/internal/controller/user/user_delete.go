package usercontroller

import (
	"context"
	"fmt"
	"net/http"
)

func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// FirebaseAuthMiddleware で設定された UserID を取得
	id, firebaseUID, err := uc.GetUserIDFromFirebaseUID(context.Background(), r)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーIDの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// Usecase層を通してユーザーを削除
	if err := uc.userUsecase.DeleteUser(context.Background(), id, firebaseUID); err != nil {
		http.Error(w, fmt.Sprintf("ユーザー削除に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// 成功レスポンス
	w.WriteHeader(http.StatusNoContent)
}
