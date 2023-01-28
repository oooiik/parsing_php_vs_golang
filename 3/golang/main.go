package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const host = "172.4.0.3"
const user = "my"
const pass = "my"
const database = "golang"

var table string = "10e" + countNol() + "-row"

var filePath string = "/app/resources/" + table + ".csv"

func countNol() string {
	if len(os.Args) < 2 {
		panic("input not found\n")
	}
	return os.Args[1]
}

func main() {

	dbTruncate()

	start := time.Now()
	defer func() {
		duration := time.Since(start)
		nolCo, err := strconv.Atoi(countNol())
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%v rows: %d µs\n", math.Pow10(nolCo), duration.Microseconds())
	}()

	// connect mysql
	db := connectMysql()
	defer db.Close()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	parser := csv.NewReader(file)

	row := 1

	per := 1000

	var lines [][]string

	for {
		line, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err.Error())
		}
		if row == 1 {
			row++
			continue
		}
		lines = append(lines, line)
		if len(lines)%per == 0 {
			dbInsert(db, lines)
			lines = [][]string{}
		}

		row++
	}
	if len(lines) > 0 {
		dbInsert(db, lines)
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

func dbInsert(db *sql.DB, values [][]string) {
	query := "INSERT INTO " + "`" + table + "`  (`uid`, `manufacturer_part_number`, `manufacturer`, `quantity`) VALUES \n"

	for key, value := range values {

		query += fmt.Sprintf("(%q, %q, %q, %q)", value[0], value[2], value[3], value[4])
		if key+1 < len(values) {
			query += ",\n"
		}
	}
	_, err := db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	err = nil
}
