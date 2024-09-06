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
	EffectiveAt time.Time
	ExpireAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
	// fmt.Printf("Document: %s inserted successfully!\n", doc.Name)
}

func MockDocuments(db *sql.DB, n int) {
	for i := 1; i <= n; i++ {
		doc := Document{
			Name:        fmt.Sprintf("Document %d", i),
			Link:        fmt.Sprintf("https://example.com/doc%d", i),
			EffectiveAt: time.Now(),
			ExpireAt:    randomAdded(time.Now()),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		InsertDocument(db, doc)
	}

	fmt.Printf("%d documents inserted successfully!\n", n)
}

func BenchmarkQuery(db *sql.DB, n int) {
	var totalElapsed1, totalElapsed2, maxElapsed1, maxElapsed2, minElapsed1, minElapsed2 time.Duration

	minElapsed1 = time.Hour
	minElapsed2 = time.Hour

	for i := 0; i < n; i++ {
		effectedDate := randomAdded(time.Now())

		start := time.Now()
		query1 := `
			SELECT * FROM documents
			WHERE ? BETWEEN effective_at AND COALESCE(expire_at, ?)`
		rows1, err := db.Query(query1, effectedDate, effectedDate.AddDate(0, 0, 1))
		if err != nil {
			log.Fatalf("Failed to execute first query: %v", err)
		}
		_ = rows1.Close()
		elapsed1 := time.Since(start)

		totalElapsed1 += elapsed1
		if elapsed1 > maxElapsed1 {
			maxElapsed1 = elapsed1
		}
		if elapsed1 < minElapsed1 {
			minElapsed1 = elapsed1
		}

		start = time.Now()
		query2 := `
			SELECT * FROM documents
			WHERE effective_at <= ? AND (expire_at IS NULL OR expire_at > ?)`
		rows2, err := db.Query(query2, effectedDate, effectedDate)
		if err != nil {
			log.Fatalf("Failed to execute second query: %v", err)
		}
		_ = rows2.Close()
		elapsed2 := time.Since(start)

		totalElapsed2 += elapsed2
		if elapsed2 > maxElapsed2 {
			maxElapsed2 = elapsed2
		}
		if elapsed2 < minElapsed2 {
			minElapsed2 = elapsed2
		}

		fmt.Printf("Query %d - BETWEEN: %v, IS NULL/OR: %v\n", i+1, elapsed1, elapsed2)
	}

	avgElapsed1 := totalElapsed1 / time.Duration(n)
	avgElapsed2 := totalElapsed2 / time.Duration(n)
	fmt.Printf("\nBETWEEN Query - Avg: %v, Max: %v, Min: %v\n", avgElapsed1, maxElapsed1, minElapsed1)
	fmt.Printf("IS NULL/OR Query - Avg: %v, Max: %v, Min: %v\n", avgElapsed2, maxElapsed2, minElapsed2)
}

func PrintSampleDocuments(db *sql.DB, n int) {
	query := `SELECT id, name, link, effective_at, expire_at, created_at, updated_at FROM documents LIMIT ?`
	rows, err := db.Query(query, n)
	if err != nil {
		log.Fatalf("Failed to fetch sample documents: %v", err)
	}
	defer rows.Close()

	fmt.Printf("Sample of %d documents:\n", n)
	for rows.Next() {
		var id int
		var name, link, effectiveAt, expireAt, createdAt, updatedAt string

		err = rows.Scan(&id, &name, &link, &effectiveAt, &expireAt, &createdAt, &updatedAt)
		if err != nil {
			log.Fatalf("Failed to scan document: %v", err)
		}
		fmt.Printf("ID: %d, Name: %s, Link: %s, EffectiveAt: %s, ExpireAt: %s, CreatedAt: %s, UpdatedAt: %s\n",
			id, name, link, effectiveAt, expireAt, createdAt, updatedAt)
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./code/go/sqlperf/db/db.sqlite3")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}
	defer db.Close()

	MockDocuments(db, 10000)
	// PrintSampleDocuments(db, 100)
	BenchmarkQuery(db, 100)
}
