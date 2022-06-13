package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "<host>:<password>@tcp(127.0.0.1:3306)/<schema name>")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}
