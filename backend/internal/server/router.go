package server

import (
	"database/sql"
	"hackathon/internal/controller"
	"github.com/gorilla/mux"
)

// ルーターを設定する関数
func NewRouter(dbConn *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// ユーザーのコントローラー設定
	userController := controller.NewUserController(dbConn)
	
	// ユーザー作成のためのエンドポイント
	router.HandleFunc("/user", userController.CreateUser).Methods("POST") // POST /user

	// ユーザー情報取得のためのエンドポイント
	router.HandleFunc("/user/{id}", userController.GetUserByID).Methods("GET") // GET /user/{id}
	
	return router
}
