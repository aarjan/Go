/*
A transaction is essentially an object that reserves a connection to the datastore.
You begin a transaction with a call to db.Begin(), and close it with a Commit() or Rollback() method on the resulting Tx variable. Under the covers, the Tx gets a connection from the pool, and reserves it for use only with that transaction. The methods on the Tx map one-for-one to methods you can call on the database itself, such as Query() and so forth.
*/
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, dbErr := sql.Open("mysql", "root:root@tcp(localhost:3306)/hello")
	checkError("dbErr", dbErr)

	stmt, stmtErr := db.Prepare("update pet set name = 'bruno',sex='M' where id = '3' and owner = 'Aarjan' ")
	checkError("stmtErr", stmtErr)

	tx, txErr := db.Begin()
	checkError("txErr", txErr)

	_, execErr := tx.Stmt(stmt).Exec()
	if execErr != nil {
		fmt.Println("doing rollback")
		tx.Rollback()
	} else {
		fmt.Println("transaction successful")
		tx.Commit()
	}

}

func checkError(s string, err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*
You should not mingle the use of transaction-related functions such as Begin() and Commit() with SQL statements such as BEGIN and COMMIT in your SQL code. Bad things might result:

The Tx objects could remain open, reserving a connection from the pool and not returning it.
The state of the database could get out of sync with the state of the Go variables representing it.
You could believe you’re executing queries on a single connection, inside of a transaction, when in reality Go has created several connections for you invisibly and some statements aren’t part of the transaction.
While you are working inside a transaction you should be careful not to make calls to the Db variable. Make all of your calls to the Tx variable that you created with db.Begin(). The Db is not in a transaction, only the Tx is. If you make further calls to db.Exec() or similar, those will happen outside the scope of your transaction, on other connections.
*/
