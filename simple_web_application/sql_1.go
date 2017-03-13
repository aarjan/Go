/*
Sample program to query database using Prepared statement
You should, in general, always prepare queries to be used multiple times. The result of preparing the query is a prepared statement, which can have placeholders (a.k.a. bind values) for parameters that you’ll provide when you execute the statement. This is much better than concatenating strings, for all the usual reasons (avoiding SQL injection attacks, for example).
*/
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, dbErr := sql.Open("mysql", "root:root@tcp(localhost:3306)/hello")
	if dbErr != nil {
		log.Fatal("dbErr", dbErr)
	}
	defer db.Close()
	stmt, stmtErr := db.Prepare("select id,name from pet where id = ?")
	if stmtErr != nil {
		log.Fatal("stmtErr ", stmtErr)
	}
	defer stmt.Close()

	rows, rowsErr := stmt.Query(2)
	defer rows.Close()

	if rowsErr != nil {
		log.Fatal("rowsErr", rowsErr)
	}

	var id int
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("rowErr", err)
		}
		log.Println(id, name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("rowsErr", err)
	}
}
