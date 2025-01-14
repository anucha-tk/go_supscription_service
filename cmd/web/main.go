package main

import (
	"database/sql"
	"log"
	"os"
	"time"
)

const webPort = "80"

func main() {
}

func initDB() *sql.DB {
	conn := connectToDB()
	if conn != nil {
		log.Panic("can't connect database")
	}
	return conn
}

func connectToDB() *sql.DB {
	counts := 0
	dns := os.Getenv("DNS")
	for {
		connection, err := openDB(dns)
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Print("connected to database")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Print("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
