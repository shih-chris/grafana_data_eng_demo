package query

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/shih-chris/grafana_data_eng_demo/pkg/types"

	_ "github.com/go-sql-driver/mysql"
)

func GetData() []types.DsTypeCount {
	// Get MySQL Env Variables
	mysqlDB := os.Getenv("MYSQL_DATABASE")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPass := os.Getenv("MYSQL_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_HOST")

	// Create MySQL Connection String
	mysqlConnectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUser, mysqlPass, mysqlHost, mysqlDB)

	// Connect to MySQL Instance
	db, err := sql.Open("mysql", mysqlConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	// SQL for count of data sources by type
	queryString := `
		select type as dsType, count(distinct id) as dsCount
		from grafana.data_source
		group by 1
		order by 2 desc
	`

	// Run SQL and create return slice combining each row
	var queryResult []types.DsTypeCount
	rows, err := db.QueryContext(ctx, queryString)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var recordType string
		var recordCount float64
		if err := rows.Scan(&recordType, &recordCount); err != nil {
			log.Fatal(err)
		}
		queryResult = append(queryResult, types.DsTypeCount{DsType: recordType, DsCount: recordCount})
	}

	return queryResult
}
