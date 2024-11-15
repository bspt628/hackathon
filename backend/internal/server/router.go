package server

import (
	"database/sql"
	"hackathon/internal/controller"
	"net/http"
)

// ルーターを設定する関数
func NewRouter(dbConn *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	// ユーザーのコントローラー設定
	userController := controller.NewUserController(dbConn)
	router.HandleFunc("/user", userController.GetUserByID)

	return router
}
