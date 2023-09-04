package main

import (
	// "errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"database/sql"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

func getData () string {
	mysqlDB := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	fmt.Println("Go Backend Start")

	mysqlConnectionString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", mysqlUser, mysqlPass, mysqlDB)

	db, err := sql.Open("mysql", mysqlConnectionString)
	if err != nil {
		panic(err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	defer db.Close()
	return "MySQL Connected Successfully"
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /metrics request\n")
	outputString := getData()
	io.WriteString(w, outputString)
}

func main() {
	http.HandleFunc("/metrics", getMetrics)

	http.ListenAndServe(":2112", nil)
}