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
	mysqlUser := os.Getenv("DB_USER")
	mysqlPwd := os.Getenv("DB_PASSWORD")
	mysqlHost := os.Getenv("DB_HOST")
	mysqlDatabase := os.Getenv("DB_NAME")

	// mysqlUser := 
	// mysqlPwd := "password"
	// mysqlHost := "tcp(23.251.145.87:3306)"
	// mysqlDatabase := "hackathon"

	// データベースを初期化します
	connStr := fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlPwd, mysqlHost, mysqlDatabase)
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
