package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"os"
	"time"
)

const host = "172.4.0.3"
const user = "my"
const pass = "my"
const database = "golang"

const table = "10e6-row"

const filePath = "/app/resources/" + table + ".csv"

func main() {
	dbTruncate()

	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Println(duration.Microseconds(), "µs")
	}()

	// connect mysql
	// 	db := connectMysql()
	// 	defer db.Close()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	parser := csv.NewReader(file)

	_, err = parser.Read()
	if err != nil {
		panic(err.Error())
	}

	for {
		_, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err.Error())
		}

		// 		dbInsert(db, line)
	}
}

func dbTruncate() {
	db := connectMysql()

	_, err := db.Query("TRUNCATE `" + table + "`")
	if err != nil {
		panic(err.Error())
	}
}

func connectMysql() *sql.DB {
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+")/"+database)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbInsert(db *sql.DB, value []string) {
	_, err := db.Exec("INSERT INTO `"+table+"`  (uid, manufacturer_part_number, manufacturer, quantity) VALUES (?, ?, ?, ?)", value[0], value[2], value[3], value[4])
	if err != nil {
		panic(err.Error())
	}
	err = nil
}
