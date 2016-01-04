package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type redirect struct {
	long_url      string
	short_url     string
	long_crc      int
	short_crc     int
	status        int
	created_by_ip int
	updated_by_ip int
	created_at    int
	updated_at    int
}

func Insert(params []map[string]interface{}) bool {
	db, err := sql.Open("mysql", "root:@/base")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	sql := "INSERT INTO redirect (long_url, short_url, long_crc, short_crc, status, created_by_ip, updated_by_ip, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?)"

	stmt, err := db.Prepare(sql)

	res, err := stmt.Exec(
		params[0]["long_url"],
		params[0]["short_url"],
		params[0]["long_crc"],
		params[0]["short_crc"],
		params[0]["status"],
		params[0]["created_by_ip"],
		params[0]["updated_by_ip"],
		params[0]["created_at"],
		params[0]["updated_at"],
	)

	id, _ := res.LastInsertId()

	fmt.Println("MMMMMMMMMM: id", id)

	return true
}
