package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DbConnection *sql.DB

func main() {
	//"user名:password値@tcp(127.0.0.1:3306)/(DB名)"
	DbConnection, _ := sql.Open("mysql", "ca_tech:ca_tech@tcp(localhost:3306)/ca_tech")

	// ユーザー情報の作成 name,token
	cmd := "INSERT INTO users (name, token) VALUES (?,?)"
	_, err := DbConnection.Exec(cmd, "tanaka", "dsfasfkjhasdjkfh")
	if err != nil {
		log.Fatalln(err)
	}
}
