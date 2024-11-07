package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

func main() {
	//
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPwd := os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := os.Getenv("MYSQL_NAME")

	connStr := fmt.Sprintf("%s:%s@(localhost:3308)/%s", mysqlUser, mysqlPwd, mysqlDatabase)
	fmt.Println(connStr)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("データベースの初期化に失敗しました: %v", err)
	}
	defer db.Close()

	fmt.Println("データベースの初期化に成功しました！")
	//fmt.Printf("%s\n", db)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)

}
