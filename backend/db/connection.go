package db

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

// InitDB はデータベース接続を初期化し、DBインスタンスを返す
func InitDB() (*sql.DB, error) {
	// DB接続のための準備
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PWD")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	
	fmt.Println(mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	

	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("データベースの初期化に失敗: %v", err)
	}

	// 接続の確認
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("データベースへの接続に失敗: %v", err)
	}

	fmt.Println("データベースの初期化に成功")
	return db, nil
}
