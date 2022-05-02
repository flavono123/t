package main

import (
	"database/sql"
	"fmt"
	"log"

	// https://go.dev/doc/tutorial/getting-started
	// Call code in an external package
	_ "github.com/go-sql-driver/mysql"
)

// mysql  Ver 15.1 Distrib 10.3.32-MariaDB, for osx10.16 (x86_64) using readline 5.1
// create table tbl (col int);
// insert into tbl values (1);
// insert into tbl values (2);
func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// DML
	result, err := db.Exec("DROP TABLE IF EXISTS tbl")
	if err != nil {
		log.Fatal(err)
	}

	n, err := result.RowsAffected()
	fmt.Printf("%d row(s) affected.\n", n)

	result, err = db.Exec("CREATE TABLE IF NOT EXISTS tbl (col int)")
	if err != nil {
		log.Fatal(err)
	}

	n, err = result.RowsAffected()
	fmt.Printf("%d row(s) affected.\n", n)

	// Prepared Statement
	stmt, err := db.Prepare("INSERT INTO tbl VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err = stmt.Exec(1)
	if err != nil {
		log.Fatal(err)
	}

	n, err = result.RowsAffected()
	fmt.Printf("%d row(s) affected.\n", n)

	result, err = stmt.Exec(2)
	if err != nil {
		log.Fatal(err)
	}

	n, err = result.RowsAffected()
	fmt.Printf("%d row(s) affected.\n", n)

	var col int
	// func (db *DB) Query(query string, args ...any) (*Rows, error)
	rows, err := db.Query("SELECT * FROM tbl")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&col)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(col)
	}

	// Transaction fail test
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO tbl VALUES (3)")
	if err != nil {
		log.Panic(err)
	}

	_, err = tx.Exec("INSERT INTO tbl VALUES ('a')")
	if err != nil {
		log.Panic(err)
	}

	err = tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}
