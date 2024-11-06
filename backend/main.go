package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//
	//mysqlUser := os.Getenv("MYSQL_USER")
	//mysqlPwd := os.Getenv("MYSQL_PWD")
	//mysqlHost := os.Getenv("MYSQL_HOST")
	//mysqlDatabase := os.Getenv("MYSQL_DATABASE")

	mysqlUser := "root"
	mysqlPwd := "rootpass"
	mysqlHost := "unix(/cloudsql/term6-hiroto-uchida:us-central1:uttc6)"
	mysqlDatabase := "hackathon"

	// データベースを初期化します
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
	fmt.Println(connStr)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("データベースの初期化に失敗しました: %v", err)
	}
	defer db.Close()

	fmt.Println("データベースの初期化に成功しました")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

}
