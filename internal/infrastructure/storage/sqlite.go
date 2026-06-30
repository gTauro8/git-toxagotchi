package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sugar_petauro/git-toxagotchi/internal/domain"
	_ "modernc.org/sqlite"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	store := &SQLiteStore{db: db}
	if err := store.migrate(); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}
	return store, nil
}

func (s *SQLiteStore) migrate() error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS pets (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			species TEXT NOT NULL,
			stage TEXT NOT NULL,
			energy INTEGER NOT NULL,
			hunger INTEGER NOT NULL,
			stress INTEGER NOT NULL,
			trust INTEGER NOT NULL,
			chaos INTEGER NOT NULL,
			experience INTEGER NOT NULL,
			age INTEGER NOT NULL,
			last_interaction_at DATETIME NOT NULL,
			mood TEXT NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			type TEXT NOT NULL,
			timestamp DATETIME NOT NULL,
			metadata TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS achievements (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			unlocked INTEGER NOT NULL DEFAULT 0,
			unlocked_at DATETIME,
			icon TEXT
		)`,
	}

	for _, q := range queries {
		if _, err := s.db.Exec(q); err != nil {
			return err
		}
	}

	// Seed achievements
	for _, ach := range domain.AllAchievements() {
		if _, err := s.db.Exec(`INSERT OR IGNORE INTO achievements (id, name, description, unlocked, icon) VALUES (?, ?, ?, 0, ?)`,
			string(ach.ID), ach.Name, ach.Description, ach.Icon); err != nil {
			return fmt.Errorf("seed achievement %s: %w", ach.ID, err)
		}
	}

	return nil
}

func (s *SQLiteStore) GetPet() (*domain.Pet, error) {
	row := s.db.QueryRow(`SELECT id, name, species, stage, energy, hunger, stress, trust, chaos, experience, age, last_interaction_at, mood FROM pets LIMIT 1`)
	p := &domain.Pet{}
	var lastInteraction string
	err := row.Scan(&p.ID, &p.Name, &p.Species, &p.Stage, &p.Energy, &p.Hunger, &p.Stress, &p.Trust, &p.Chaos, &p.Experience, &p.Age, &lastInteraction, &p.Mood)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	p.LastInteractionAt, _ = time.Parse(time.RFC3339, lastInteraction)
	return p, nil
}

func (s *SQLiteStore) SavePet(p *domain.Pet) error {
	_, err := s.db.Exec(`INSERT OR REPLACE INTO pets (id, name, species, stage, energy, hunger, stress, trust, chaos, experience, age, last_interaction_at, mood)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		p.ID, p.Name, p.Species, string(p.Stage), p.Energy, p.Hunger, p.Stress, p.Trust, p.Chaos, p.Experience, p.Age,
		p.LastInteractionAt.Format(time.RFC3339), string(p.Mood))
	return err
}

func (s *SQLiteStore) SaveEvent(e *domain.Event) error {
	meta, _ := json.Marshal(e.Metadata)
	_, err := s.db.Exec(`INSERT INTO events (type, timestamp, metadata) VALUES (?, ?, ?)`,
		string(e.Type), e.Timestamp.Format(time.RFC3339), string(meta))
	return err
}

func (s *SQLiteStore) GetRecentEvents(limit int) ([]*domain.Event, error) {
	rows, err := s.db.Query(`SELECT type, timestamp, metadata FROM events ORDER BY timestamp DESC LIMIT ?`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*domain.Event
	for rows.Next() {
		e := &domain.Event{}
		var ts, meta string
		if err := rows.Scan(&e.Type, &ts, &meta); err != nil {
			continue
		}
		e.Timestamp, _ = time.Parse(time.RFC3339, ts)
		_ = json.Unmarshal([]byte(meta), &e.Metadata)
		events = append(events, e)
	}
	return events, nil
}

func (s *SQLiteStore) GetAchievements() ([]domain.Achievement, error) {
	rows, err := s.db.Query(`SELECT id, name, description, unlocked, unlocked_at, icon FROM achievements`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var achs []domain.Achievement
	for rows.Next() {
		var a domain.Achievement
		var unlockedAt sql.NullString
		var unlocked int
		if err := rows.Scan(&a.ID, &a.Name, &a.Description, &unlocked, &unlockedAt, &a.Icon); err != nil {
			continue
		}
		a.Unlocked = unlocked == 1
		if unlockedAt.Valid {
			t, _ := time.Parse(time.RFC3339, unlockedAt.String)
			a.UnlockedAt = &t
		}
		achs = append(achs, a)
	}
	return achs, nil
}

func (s *SQLiteStore) Close() error {
	return s.db.Close()
}
