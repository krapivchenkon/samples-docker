package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	// "time"
)

var (
	DB_USER     = "docker"
	DB_PASSWORD = "docker"
	DB_NAME     = "docker"
	DB_HOST     = "localhost"
	DB_PORT     = os.Getenv("PG_PORT")
)

func main() {
	// dbinfo := fmt.Sprintf("user=%s password=%s host=localhost port=%s dbname=%s sslmode=disable",
	// DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	dbinfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	fmt.Printf("Connecting to %s\n", dbinfo)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT id, name,address FROM COMPANY")
	checkErr(err)
	defer rows.Close()
	fmt.Println("uid | username | address ")
	for rows.Next() {
		var id int
		var name string
		var address string

		err = rows.Scan(&id, &name, &address)
		checkErr(err)
		fmt.Printf("%3v | %8v | %6v \n", id, name, address)
	}
	err = rows.Err()
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
