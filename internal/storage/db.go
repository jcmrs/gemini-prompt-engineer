package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the SQLite database at the given path.
func InitDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	log.Println("Database initialized successfully.")
	return db, nil
}

func createTables(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS conversations (
			id TEXT PRIMARY KEY,
			title TEXT,
			created_at TEXT,
			updated_at TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS messages (
			id TEXT PRIMARY KEY,
			conversation_id TEXT,
			role TEXT,
			content TEXT,
			attachments JSON,
			model TEXT,
			model_config JSON,
			prompt_version TEXT,
			created_at TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS prompts (
			id TEXT PRIMARY KEY,
			version TEXT,
			author TEXT,
			intent TEXT,
			description TEXT,
			prompt_text TEXT,
			settings JSON,
			examples JSON,
			created_at TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS runs (
			id TEXT PRIMARY KEY,
			prompt_id TEXT,
			model TEXT,
			settings JSON,
			status TEXT,
			created_at TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS evaluations (
			id TEXT PRIMARY KEY,
			run_id TEXT,
			metrics JSON,
			created_at TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS attachments (
			id TEXT PRIMARY KEY,
			filename TEXT,
			path TEXT,
			mimetype TEXT,
			text_extract TEXT,
			created_at TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS audit_logs (
			id TEXT PRIMARY KEY,
			action TEXT,
			actor TEXT,
			details JSON,
			created_at TEXT
		)`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS messages_fts USING fts5(content, tokenize = 'porter')`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS prompts_fts USING fts5(prompt_text, tokenize = 'porter')`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}
