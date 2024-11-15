package controller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
	// ぱすぱらめーたからユーザーIDを取得
	vars := mux.Vars(r)
	userID := vars["id"]
	// クエリパラメータからユーザーIDを取得
	// userID := r.URL.Query().Get("id")
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

// CreateUser は新規ユーザーを作成するエンドポイント
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// リクエストボディからユーザー情報を取得
	var request struct {
		Email        string `json:"email"`
		PasswordHash string `json:"password_hash"`
		Username     string `json:"username"`
		DisplayName  string `json:"display_name"`
	}

	// リクエストのJSONデータを構造体にバインド
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("リクエストの解析に失敗しました: %v", err), http.StatusBadRequest)
		return
	}

	// 必須フィールドのバリデーション
	if request.Email == "" || request.PasswordHash == "" || request.Username == "" {
		http.Error(w, "必須フィールドが不足しています", http.StatusBadRequest)
		return
	}

	// 新規ユーザーを作成
	user, err := uc.userUsecase.CreateUser(context.Background(), request.Email, request.PasswordHash, request.Username, request.DisplayName)
	if err != nil {
		http.Error(w, fmt.Sprintf("ユーザー作成に失敗しました: %v", err), http.StatusInternalServerError)
		return
	}

	// JSON形式でレスポンスを返す
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
	}
}