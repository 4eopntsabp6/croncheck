// Package history tracks previously validated cron expressions
// and provides retrieval and persistence utilities.
package history

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// Entry represents a single history record for a cron expression.
type Entry struct {
	Expression string    `json:"expression"`
	Description string  `json:"description"`
	Valid       bool     `json:"valid"`
	CheckedAt   time.Time `json:"checked_at"`
}

// History holds a list of cron expression entries.
type History struct {
	Entries []Entry `json:"entries"`
}

// Add appends a new entry to the history.
func (h *History) Add(expr, description string, valid bool) {
	h.Entries = append(h.Entries, Entry{
		Expression:  expr,
		Description: description,
		Valid:       valid,
		CheckedAt:   time.Now().UTC(),
	})
}

// Last returns the most recently added entry, or nil if empty.
func (h *History) Last() *Entry {
	if len(h.Entries) == 0 {
		return nil
	}
	return &h.Entries[len(h.Entries)-1]
}

// Load reads history from a JSON file at the given path.
func Load(path string) (*History, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &History{}, nil
	}
	if err != nil {
		return nil, err
	}
	var h History
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}
	return &h, nil
}

// Save writes the history to a JSON file at the given path.
func Save(path string, h *History) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// DefaultPath returns the default path for the history file.
func DefaultPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".croncheck_history.json"
	}
	return filepath.Join(home, ".croncheck", "history.json")
}
