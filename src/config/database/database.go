package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"time"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

// Database contains a *sql.DB pointer and a migration runner
type Database struct {
	DB              *sql.DB
	migrationRunner MigrationRunner
}

// NewDatabase returns a new Database instance and invokes
// its migration runner
func NewDatabase() Database {
	database := newDB()

	migrationRunner := NewMigrationRunner(database)
	migrationRunner.ProcessMigrations()

	return Database{DB: database, migrationRunner: migrationRunner}
}

type databaseSecrets struct {
	Port     uint16 `json:"db_port"`
	Host     string `json:"db_host"`
	User     string `json:"db_user"`
	Password string `json:"db_password"`
	Name     string `json:"db_name"`
	SSLMode  string `json:"db_sslmode"`
}

func loadSecrets() databaseSecrets {
	path := filepath.FromSlash("resources/properties/secrets.json")
	log.Printf("Reading properties from %s \n", path)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}

	var secrets databaseSecrets

	err = json.Unmarshal(content, &secrets)
	if err != nil {
		log.Panic(err)
	}

	return secrets
}

func ping(db *sql.DB, retries uint8) {
	success := false

	for i := uint8(0); i < retries && !success; i++ {
		if err := db.Ping(); err != nil {
			log.Printf("Failed to ping database, retrying in 500 ms: %s", err.Error())
			time.Sleep(500 * time.Millisecond)
		} else {
			success = true
		}
	}

	if !success {
		log.Panicf("Failed to successfully ping database after %d retries", retries)
	}
}

func newDB() *sql.DB {
	secrets := loadSecrets()
	connectionString := fmt.Sprintf(
		"port=%d host=%s user=%s password=%s dbname=%s sslmode=%s",
		secrets.Port, secrets.Host, secrets.User, secrets.Password, secrets.Name, secrets.SSLMode)

	log.Printf("Connecting to database %s:%d \n", secrets.Host, secrets.Port)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	ping(db, 30)

	log.Printf("Successfully connected to database %s:%d \n", secrets.Host, secrets.Port)

	return db
}
