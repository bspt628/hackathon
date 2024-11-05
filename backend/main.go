package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

    mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPwd := os.Getenv("MYSQL_PWD")
    mysqlHost := os.Getenv("MYSQL_HOST")
    mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	// データベースを初期化します
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
    db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("データベースの初期化に失敗しました: %v", err)
	}
	defer db.Close()

	fmt.Println("データベースの初期化に成功しました")
}
