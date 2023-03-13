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

const maxGoroutinesOnlyTime = 10

var parser *csv.Reader
var dbConnect *sql.DB

func countNol() string {
	if len(os.Args) < 2 {
		panic("input not found\n")
	}
	return os.Args[1]
}

func main() {

	dbTruncate()

	//Timer
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

	// channel
	lines := make(chan []string, maxGoroutinesOnlyTime)

	go reader(lines)
	writer(lines)
}

func reader(lines chan []string) {
	for {
		line, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err.Error())
		}
		lines <- line
	}
}

func writer(lines chan []string) {
	for {
		select {
		case value, openChan := <-lines:
			if !openChan {
				break
			}
			writerLine(value)
		}
	}
}

func writerLine(value []string) {
	dbInsert(value)
}

func getParser() *csv.Reader {
	if parser == nil {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}

		parser = csv.NewReader(file)

		_, err = parser.Read() // 1 line
		if err != nil {
			panic(err.Error())
		}
	}

	return parser
}

func read() ([]string, bool) {
	line, err := getParser().Read()
	if err == io.EOF {
		return nil, false
	}
	if err != nil {
		panic(err.Error())
	}
	return line, true
}

func dbInsert(value []string) bool {
	db := connectMysql()
	_, err := db.Exec("INSERT INTO "+"`"+table+"`  (uid, manufacturer_part_number, manufacturer, quantity) VALUES (?, ?, ?, ?)", value[0], value[2], value[3], value[4])
	if err != nil {
		panic(err.Error())
	}
	err = nil
	return true
}

func dbTruncate() {
	db := connectMysql()

	_, err := db.Query("TRUNCATE `" + table + "`")
	if err != nil {
		panic(err.Error())
	}
}

func connectMysql() *sql.DB {
	if dbConnect == nil {
		db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+")/"+database)
		if err != nil {
			panic(err.Error())
		}
		dbConnect = db
	}
	return dbConnect
}
