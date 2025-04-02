package database

import (
    "database/sql"
    "log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	instance *sql.DB
	once     sync.Once
)

func Connect() (*sql.DB, error){
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}

		// Configuration max connection 
		db.SetMaxOpenConns(25)  
		db.SetMaxIdleConns(25)  
		db.SetConnMaxLifetime(1 * time.Minute)  

		err = db.Ping()
		if err != nil {
            log.Fatalf("Error pinging database: %v", err)
        }

		instance = db
		log.Println("Connected to database successfully")

	})

	return instance, nil
}