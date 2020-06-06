package database

import (
	"database/sql"
	"log"
	"time"
)

type migration struct {
	Version             string
	Description         string
	Script              string
	Checksum            string
	ScriptExecutionTime time.Time
}

// MigrationRunner reads migration files and executes them
type MigrationRunner struct {
	database *sql.DB
}

// NewMigrationRunner creates the migration table if it does not
// exist, and returns a new MigrationRunner instance
func NewMigrationRunner(database *sql.DB) MigrationRunner {
	mr := MigrationRunner{database: database}
	mr.ensureVersionTableExistence()
	return mr
}

func (mr *MigrationRunner) databaseVersionTableExists() bool {
	var result bool

	rows, err := mr.database.Query(`
	SELECT EXISTS (
		SELECT 
		FROM 
			information_schema.tables 
		WHERE  
			table_schema = 'public'
		AND    
			table_name   = 'database_version'
	);`)

	if err != nil {
		log.Fatal(err)
	}

	if rows.Next() {
		rows.Scan(&result)
	}

	return result
}

func (mr *MigrationRunner) ensureVersionTableExistence() {
	if mr.databaseVersionTableExists() {
		return
	}

	log.Println("Creating database_version table")
	_, err := mr.database.Exec(
		`CREATE TABLE public.database_version(
			version character varying(50),
			description character varying(200),
			script TEXT,
			checksum character varying(50),
			script_execution_time timestamp without time zone NOT NULL DEFAULT now(),
			CONSTRAINT database_version_pk PRIMARY KEY (version)
		)`)

	if err != nil {
		log.Fatal(err)
	}
}
