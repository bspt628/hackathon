package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"context"
	"encoding/json"
	"hackathon/db/sqlc/generated"
	"hackathon/internal/dao"
	"hackathon/internal/usecase"
)

func main() {
	// .env ファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env ファイルの読み込みに失敗しました: %v", err)
	}

	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_NAME")

	connStr := fmt.Sprintf("%s:%s@(localhost:3308)/%s", mysqlUser, mysqlPwd, mysqlDatabase)
	fmt.Println(connStr)
	db0, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("データベースの初期化に失敗しました: %v", err)
	}
	defer db0.Close()

	fmt.Println("データベースの初期化に成功しました！")

	if err := db0.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("pingに成功しました！")

	// Queries, DAO, Usecaseのインスタンスを作成
	queries := db.New(db0)
	userDAO := dao.NewUserDAO(queries)
	userUsecase := usecase.NewUserUsecase(userDAO)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ユーザーIDを指定してユーザーを取得するエンドポイント
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		// クエリパラメータからユーザーIDを取得
		userID := r.URL.Query().Get("id")
		if userID == "" {
			http.Error(w, "IDパラメータが指定されていません", http.StatusBadRequest)
			return
		}

		// コンテキストとともにユーザーを取得
		user, err := userUsecase.GetUserByID(context.Background(), userID)
		if err != nil {
			http.Error(w, fmt.Sprintf("ユーザーの取得に失敗しました: %v", err), http.StatusInternalServerError)
			return
		}

		// JSON形式でレスポンスを返す
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, fmt.Sprintf("レスポンスのエンコードに失敗しました: %v", err), http.StatusInternalServerError)
		}
	})

	// サーバーを起動
	log.Printf("サーバーをポート%sで起動します...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
