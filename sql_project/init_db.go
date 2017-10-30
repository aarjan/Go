package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"time"
)

func init() {
	// get data from config file
	buf, err := ioutil.ReadFile("config.json")
	checkError("ReadFile error ", err)
	err = json.Unmarshal(buf, &c)
	checkError("Unmarshal error ", err)

	// Open database connection pool
	db, err = sql.Open("mysql", c.Username+":"+c.Password+"@tcp(localhost:"+c.Port+")/")
	checkError("open connection error ", err)

	// Check if database exists!, else create a new one
	_, err = db.Exec("create database if not exists " + c.Dbname)
	checkError("create db error ", err)
	defer db.Close()

	// Use that database
	_, err = db.Exec("use " + c.Dbname)

	// tableName := strings.Join(strings.Fields(fmt.Sprint(time.Now().Clock())), "_") + "_ReportName"
	_, err = db.Exec("Create table ReportName (id int primary key not null, year varchar(50) not null)")
	if err != nil {
		log.Fatal("Could not create table,err", err)
	}

	// Create dummy values
	for i := 1; i <= 5; i++ {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		a := random.Perm(3)
		year := strings.Join(strings.Split(fmt.Sprint(time.Now().AddDate(a[0], a[1], a[2]).Date()), " "), "-")

		stmt, er := db.Prepare("insert into ReportName values(?,?)")
		checkError("error Preparing query, ", er)

		defer stmt.Close()
		_, err = stmt.Exec(i, year)
		checkError("error executing query, ", err)
	}

	log.Printf("DB initialized with database,%s and table, ReportName",c.Dbname)

}
