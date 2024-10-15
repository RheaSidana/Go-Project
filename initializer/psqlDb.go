package initializer

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
)
var Db *sql.DB

func ConnectToDB() {
	dbUrl := getDbUrl()
	// fmt.Println("DB URL : ", dbUrl)
	var err error
	Db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	if err := Db.Ping(); err != nil {
		log.Fatalf("error pinging database: %v", err)
	}

	log.Println("Connected to the database successfully!")

}

func CloseDb(){
	Db.Close()
}
func getDbUrl() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE") 

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, 
		dbName, dbSSLMode,
	)
}