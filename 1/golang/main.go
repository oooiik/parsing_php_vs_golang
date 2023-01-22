package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/gosuri/uilive"
	"os"
	"time"
)

const host = "mysql"
const user = "my"
const pass = "my"
const database = "golang"

const table = "10e5-row"

const filePath = "/app/resources/" + table + ".csv"

func main() {
	defer dbTruncate()

	start := time.Now()
	defer func() {
		duration := time.Since(start)
		fmt.Println(duration.Microseconds(), "µs")
	}()

	//fmt.Println("Start!")
	//defer fmt.Println("Finish!")

	// connect mysql
	db := connectMysql(host, user, pass, database)
	defer db.Close()

	// open file
	file := openFile(filePath)
	defer file.Close()

	//writer := uilive.New()
	//writer.Start()
	//defer writer.Stop()

	for i, value := range readCsvFile(file) {
		if i == 0 {
			continue
		}
		dbInsert(db, table, value)
		//fmt.Fprintf(writer, "Row line - %d \n", i)
	}

}

func connectMysql(host, user, pass, database string) *sql.DB {
	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+")/"+database)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func dbInsert(db *sql.DB, table string, value []string) {
	_, err := db.Query("INSERT INTO `"+table+"`  (uid, manufacturer_part_number, manufacturer, quantity) VALUES (?, ?, ?, ?)", value[0], value[2], value[3], value[4])
	if err != nil {
		panic(err.Error())
	}
}

func dbTruncate() {
	db := connectMysql(host, user, pass, database)

	_, err := db.Query("TRUNCATE `" + table + "`")
	if err != nil {
		panic(err.Error())
	}
}

func openFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err.Error())
	}
	return file
}

func readCsvFile(file *os.File) [][]string {
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		panic(err.Error())
	}
	return records
}
