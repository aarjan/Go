package main

import (
	"database/sql"
	"fmt"
	"log"

	"flag"
	"strings"
	"time"
)

var db *sql.DB
var c = config{}

// Contents of config file
type config struct {
	Username string
	Password string
	Dbname   string
	Port     string
}

func checkError(s string, err error) {

	if err != nil {
		log.Fatal(s, err)
	}

}

func main() {

	table := flag.String("tableName", "null", "select table")
	pattern := flag.Bool("change_pattern", false, "change pattern of year to DD/MM/YYYY")
	flag.Parse()

	if *pattern {
		tableName := strings.Join(strings.Fields(fmt.Sprint(time.Now().Clock())), ".")

		// Create new table
		_, err := db.Exec("Create table " + tableName + "_ReportName (ID primary key not null, year Date)")
		checkError(" change error", err)

		var id string
		var year string

		// Get the current dates
		rows, err := db.Query("select id,year from " + *table)
		checkError("get year error", err)
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&id, &year)
			checkError("scan year error", err)

			// Refactor the data
			d := strings.Split(year, "/")
			refYear := d[2] + "/" + d[1] + "/" + d[0]

			// Refactor the Field 'Year'
			_, err = db.Exec("insert into " + tableName + " values (" + id + "," + refYear + ")")
			if err != nil {
				fmt.Println("Success")
			}


		}
	}
}
