package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

var cfg = pq.Config{
	Host:           "localhost",
	Port:           5432,
	User:           "user",
	ConnectTimeout: 5 * time.Second,
	Password:       "0088",
	SSLMode:        pq.SSLModeDisable,
	Database:       "go",
}

func main() {
	fmt.Println("welcome to chai")

	c, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Create connection pool.
	db := sql.OpenDB(c)
	defer db.Close()

	// Make sure it works.
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// 	rows, err := db.Query(`CREATE TABLE IF NOT EXISTS public.students (
	//     student_id SERIAL PRIMARY KEY,
	//     first_name VARCHAR(50) NOT NULL,
	//     last_name VARCHAR(50),
	//     email VARCHAR(100) UNIQUE
	// );`)

	rows, err := db.Query(`SELECT student_id, first_name FROM students;`)

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

}
