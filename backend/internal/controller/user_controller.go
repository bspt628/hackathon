package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/dao"
	"hackathon/internal/usecase"
	"net/http"
)

type UserController struct {
    userUsecase *usecase.UserUsecase // Usecaseへの依存性注入
}

func NewUserController(dbConn *sql.DB) *UserController {
	queries := db.New(dbConn)
	userDAO := dao.NewUserDAO(queries)
	userUsecase := usecase.NewUserUsecase(userDAO)
	return &UserController{userUsecase: userUsecase}
}

// GetUserByID はユーザーIDを指定してユーザー情報を取得するエンドポイント
func (uc *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータからユーザーIDを取得
	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.Error(w, "IDパラメータが指定されていません", http.StatusBadRequest)
		return
	}

	// コンテキストとともにユーザーを取得
	user, err := uc.userUsecase.GetUserByID(context.Background(), userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザーの取得に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// JSON形式でレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}