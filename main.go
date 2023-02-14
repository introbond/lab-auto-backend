package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	connString := `server=192.168.11.63;
		user id=sa;
		password=P@ssw0rd;
		database=backend;
		encrypt=disable;
		trustServerCertificate=true;`

	db, err := sql.Open("mssql", connString)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	var name string
	err = db.QueryRow(`SELECT name FROM test WHERE id = $1`, 2).Scan(&name)
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}

	fmt.Println("Name:", name)
}
