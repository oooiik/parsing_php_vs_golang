package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
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

	fmt.Println("Start!!!")
	defer fmt.Println("Finish!!!")
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		nolCo, err := strconv.Atoi(countNol())
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("%v rows: %d µs\n", math.Pow10(nolCo), duration.Microseconds())
	}()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	parser := csv.NewReader(file)

	multiPer := 100
	per := 100

	lineAll, err := parser.ReadAll()
	if err != nil {
		panic(err.Error())
	}
	linesCount := len(lineAll)

	var wg sync.WaitGroup
	for i := 1; i < linesCount; i += multiPer * per {
		wg.Add(1)
		if i+multiPer*per < linesCount {
			go multi(per, lineAll[i:i+multiPer*per], &wg)
		} else {
			go multi(per, lineAll[i:], &wg)
		}
	}
	wg.Wait()
}

var multiCount int = 0

func multi(per int, lines [][]string, wg *sync.WaitGroup) {
	defer wg.Done()

	multiCount++
	son := multiCount
	fmt.Println(son, "- Start")
	defer fmt.Println(son, "- Finish")

	// connect mysql
	db := connectMysql()
	defer db.Close()

	for i := 1; i < len(lines); i += per {
		if i+per < len(lines) {
			dbInsert(db, lines[i:i+per])
		} else {
			dbInsert(db, lines[i:])
		}
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
