// Sample program to query the database and retrieve results
package main

import (
	"database/sql"
	// "_" is used to initialize the driver only
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// opening connection through mysql driver
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var (
		name  string
		sex   string
		birth string
	)

	// Quering the database
	rows, er := db.Query("select name,sex,birth from pet where sex = ? and name = ? order by birth", "F", "Pumpu")
	if er != nil {
		log.Fatal(er)
	}
	defer rows.Close()

	/*
	   As long as there’s an open result set (represented by rows), the underlying connection is busy and can’t be used for any other query.
	   That means it’s not available in the connection pool.
	   If you iterate over all of the rows with rows.Next(), eventually you’ll read the last row,
	   and rows.Next() will encounter an internal EOF error and call rows.Close() for you.
	   But if for some reason you exit that loop – an early return, or so on – then the rows doesn’t get closed, and the connection remains open.
	   (It is auto-closed if rows.Next() returns false due to an error, though).
	   This is an easy way to run out of resources.
	*/
	for rows.Next() {
		err = rows.Scan(&name, &sex, &birth)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name, sex, birth)
	}
	err = rows.Err() //You should always check for an error at the end of the for rows.Next() loop. If there’s an error during the loop, you need to know about it
	if err != nil {
		log.Fatal(err)
	}
}
