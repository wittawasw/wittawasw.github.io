package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Document struct {
	Name        string
	Link        string
	EffectiveAt string
	ExpireAt    string
	CreatedAt   string
	UpdatedAt   string
}

// Randomly add number of days to time, can be minus to
// make it into the past in range +/- 365 days
func randomAdded(t time.Time) time.Time {
	direction := rand.Intn(2) == 0
	randDays := rand.Intn(365) + 1
	if direction {
		return t.AddDate(0, 0, randDays)
	}

	return t.AddDate(0, 0, -randDays)
}

func InsertDocument(db *sql.DB, doc Document) {
	insertSQL := `
		INSERT INTO documents
		(name, link, effective_at, expire_at, created_at, updated_at)
		VALUES
		(?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(insertSQL, doc.Name, doc.Link, doc.EffectiveAt, doc.ExpireAt, doc.CreatedAt, doc.UpdatedAt)

	if err != nil {
		log.Fatalf("Failed to insert document: %v", err)
	}
	fmt.Printf("Document: %s inserted successfully!\n", doc.Name)
}

func MockDocuments(db *sql.DB, n int) {
	for i := 1; i <= n; i++ {
		doc := Document{
			Name:        fmt.Sprintf("Document %d", i),
			Link:        fmt.Sprintf("https://example.com/doc%d", i),
			EffectiveAt: time.Now().Format(time.RFC3339),
			ExpireAt:    randomAdded(time.Now()).Format(time.RFC3339),
			CreatedAt:   time.Now().Format(time.RFC3339),
			UpdatedAt:   time.Now().Format(time.RFC3339),
		}
		InsertDocument(db, doc)
	}

	fmt.Printf("%d documents inserted successfully!\n", n)
}

func main() {
	db, err := sql.Open("sqlite3", "./code/go/sqlperf/db/db.sqlite3")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}
	defer db.Close()

	MockDocuments(db, 1000)
}
