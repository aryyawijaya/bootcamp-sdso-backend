package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // don't forget to add it. It doesn't be added automatically
	_ "github.com/jackc/pgx/v5/stdlib"
)

var Db *sql.DB //created outside to make it global.

// make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() {

	host := os.Getenv("POSTGRES_HOST")
	// port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT")) // don't forget to convert int since port is int type.
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	pass := os.Getenv("POSTGRES_PASSWORD")

	// fmt.Printf("host=%s\nport=%d\nuser=%s\ndbname=%s\npass=%s", host, port, user, dbname, pass)

	// set up postgres sql to open it.
	// psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
	// 	host, port, user, dbname, pass)
	psqlSetup := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s",
		host, user, pass, port, dbname)

	db, errSql := sql.Open("pgx", psqlSetup)
	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", errSql)
		panic(errSql)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}
}
