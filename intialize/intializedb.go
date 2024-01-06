package intialize

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

// Connect to DB
type PostgreDB struct {
	DB *sql.DB
}

func InitDB() (PostgreDB, error) {
	fmt.Println("Init DB connection - Postgres")
	var store = PostgreDB{}
	var err error
	store.DB, err = sql.Open("postgres", "postgresql://apiuser:apiuser@192.168.86.153/postgres?sslmode=disable")
	if err != nil {
		fmt.Println("Error opening DB connection")
		log.Fatal(err.Error())
	}
	fmt.Println("Connected to DB!!")
	return store, err
}
