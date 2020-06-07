package database

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	//PostgreSQL driver
	_ "github.com/lib/pq"
)

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
