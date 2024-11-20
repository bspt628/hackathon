package db

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

// InitDB はデータベース接続を初期化し、DBインスタンスを返す
func InitDB() (*sql.DB, error) {
	// 環境変数からデータベース接続情報を取得
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_NAME")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")

	// 接続文字列を作成
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
		mysqlUser, mysqlPwd, mysqlHost, mysqlPort, mysqlDatabase)
	fmt.Println(connStr)
	// データベース接続を開く
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("データベースの初期化に失敗しました: %v", err)
	}

	// 接続の確認
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("データベースへの接続に失敗しました: %v", err)
	}

	fmt.Println("データベースの初期化に成功しました！")
	return db, nil
}
