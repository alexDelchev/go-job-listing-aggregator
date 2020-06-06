package database

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

func (mr *MigrationRunner) getMigration(version string) migration {
	var result migration

	query := `
		SELECT
			version,
			description,
			script,
			checksum,
			script_execution_time
		FROM
			public.database_version
		WHERE
			version = $1`

	rows, err := mr.database.Query(query, version)
	if err != nil {
		log.Fatal(err)
	}

	if rows.Next() {
		rows.Scan(&result.Version, &result.Description,
			&result.Script, &result.Checksum,
			&result.ScriptExecutionTime)
	}

	return result
}

func (mr *MigrationRunner) persistAppliedMigration(migration *migration) {
	statement := `
		INSERT INTO public.database_version(
			version, 
			description, 
			script, 
			checksum
		) VALUES (
			$1,
			$2,
			$3,
			$4
		)`

	if _, err := mr.database.Exec(statement,
		migration.Version, migration.Description, migration.Script, migration.Checksum); err != nil {
		log.Fatalf("Failed persisting applied migration %s - %s: %v",
			migration.Version, migration.Description, err)
	}
}

func loadMigrationFilePaths() ([]string, error) {
	var files []string

	migrationDirPath := filepath.FromSlash("resources/db/migration/")

	err := filepath.Walk(migrationDirPath, func(migrationDirPath string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, migrationDirPath)
		}
		return nil
	})

	return files, err
}

func generateChecksum(data []byte) string {
	hash := md5.New()

	if _, err := io.Copy(hash, bytes.NewReader(data)); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func getMigrationName(path string) string {
	tokens := strings.Split(path, string(os.PathSeparator))
	return tokens[len(tokens)-1]
}

func loadMigrationFile(path string) migration {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Panic(err)
	}

	migrationName := getMigrationName(path)

	migrationFileNameRegex := regexp.MustCompile("(?i)(.*)__(.*)\\.(sql|txt)")

	version := migrationFileNameRegex.ReplaceAll([]byte(migrationName), []byte("$1"))

	description := migrationFileNameRegex.ReplaceAll([]byte(migrationName), []byte("$2"))

	checksum := generateChecksum(content)

	return migration{
		Version:     string(version),
		Description: string(description),
		Script:      string(content),
		Checksum:    checksum}
}

func (mr *MigrationRunner) isMigrationProcessed(migration *migration) bool {
	storedMigration := mr.getMigration(migration.Version)

	if storedMigration.Version == "" {
		return false
	}

	if storedMigration.Checksum != migration.Checksum {
		log.Fatalf("Validation failed for migration %s: Expected checksum was %s, but got %s",
			migration.Version, storedMigration.Checksum, migration.Checksum)
	}

	return true
}

func (mr *MigrationRunner) applyMigration(migration *migration) {
	log.Printf("Applying migration %s - %s", migration.Version, migration.Description)

	if _, err := mr.database.Exec(migration.Script); err != nil {
		log.Fatalf("Failed applying %s - %s: %v", migration.Version, migration.Description, err)
	}

	mr.persistAppliedMigration(migration)
	log.Printf("Successfully applied migration %s - %s", migration.Version, migration.Description)
}

// ProcessMigrations prints the migrations files
func (mr *MigrationRunner) ProcessMigrations() {
	migrationFiles, err := loadMigrationFilePaths()
	if err != nil {
		log.Panic(err)
	}

	for _, path := range migrationFiles {
		migration := loadMigrationFile(path)

		if mr.isMigrationProcessed(&migration) {
			continue
		}

		mr.applyMigration(&migration)
	}

	log.Printf("Successfully processed %d migrations \n", len(migrationFiles))
}
