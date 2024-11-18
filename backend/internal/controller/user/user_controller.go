package controller

import (
	"database/sql"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/dao/user"
	"hackathon/internal/usecase/user"
	"hackathon/internal/auth"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct {
    userUsecase *usecase.UserUsecase // Usecaseへの依存性注入
}

func NewUserController(dbConn *sql.DB) *UserController {
	queries := sqlc.New(dbConn)
	userDAO := dao.NewUserDAO(queries)
	userUsecase := usecase.NewUserUsecase(userDAO)
	return &UserController{userUsecase: userUsecase}
}

// ユーザー情報取得のハンドラー
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Header.Get("UserID") // Firebase認証から設定したUserID

    // ここでuserIDを使ってユーザー情報を取得する
    userRecord, err := auth.GetUserInfo(userID)
    if err != nil {
        http.Error(w, fmt.Sprintf("User retrieval failed: %v", err), http.StatusInternalServerError)
        return
    }

    // ユーザー情報をJSONとして返す
    response := map[string]string{"userID": userRecord.UID, "email": userRecord.Email}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}