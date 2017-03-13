// Sample program to modifying data into database
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, dbErr := sql.Open("mysql", "root:root@tcp(localhost:3306)/hello")
	checkError("dbErr", dbErr)

	stmt, stmtErr := db.Prepare("insert into pet values(NULL,'fuchhi','sailu','F',null) ")
	checkError("stmtErr", stmtErr)

	// Executing a statement produces a sql.result ; that gives access to statement metadata.
	res, resErr := stmt.Exec()
	checkError("resErr", resErr)

	lastID, IDErr := res.LastInsertId()
	checkError("IDErr", IDErr)

	rowCnt, err := res.RowsAffected()
	checkError("rowErr", err)

	log.Printf("ID = %d, affected = %d \n", lastID, rowCnt)

	// If we want to ignore the result of prepared statement
	_, updErr := db.Exec("update pet set id = ? where id= ?", 11, 14)
	checkError("updErr", updErr)

}

func checkError(s string, err error) {
	if err != nil {
		log.Fatal(s, err)
	}
}

/*
_,err := db.Exec("Delete from pet") // OK
_,err := db.Query("Delete from pet") // BAD

They do not do the same thing, and you should never use Query() like this. The Query() will return a sql.Rows, which reserves a database connection until the sql.Rows is closed. Since there might be unread data (e.g. more data rows), the connection can not be used. In the example above, the connection will never be released again. The garbage collector will eventually close the underlying net.Conn for you, but this might take a long time. Moreover the database/sql package keeps tracking the connection in its pool, hoping that you release it at some point, so that the connection can be used again. This anti-pattern is therefore a good way to run out of resources (too many connections, for example).

*/
