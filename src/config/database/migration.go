package database

import (
	"time"
)

type migration struct {
	Version             string
	Description         string
	Script              string
	Checksum            string
	ScriptExecutionTime time.Time
}
