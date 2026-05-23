package history_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/user/croncheck/internal/history"
)

func TestAdd_IncreasesLength(t *testing.T) {
	h := &history.History{}
	h.Add("* * * * *", "every minute", true)
	if len(h.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(h.Entries))
	}
}

func TestLast_ReturnsNilWhenEmpty(t *testing.T) {
	h := &history.History{}
	if h.Last() != nil {
		t.Fatal("expected nil for empty history")
	}
}

func TestLast_ReturnsMostRecent(t *testing.T) {
	h := &history.History{}
	h.Add("* * * * *", "every minute", true)
	h.Add("0 * * * *", "hourly", true)
	last := h.Last()
	if last == nil {
		t.Fatal("expected non-nil last entry")
	}
	if last.Expression != "0 * * * *" {
		t.Errorf("expected '0 * * * *', got '%s'", last.Expression)
	}
}

func TestSaveAndLoad_RoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "history.json")

	h := &history.History{}
	h.Add("* * * * *", "every minute", true)
	h.Add("bad expr", "", false)

	if err := history.Save(path, h); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := history.Load(path)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if len(loaded.Entries) != 2 {
		t.Errorf("expected 2 entries, got %d", len(loaded.Entries))
	}
	if loaded.Entries[0].Expression != "* * * * *" {
		t.Errorf("unexpected expression: %s", loaded.Entries[0].Expression)
	}
}

func TestLoad_MissingFileReturnsEmpty(t *testing.T) {
	h, err := history.Load("/nonexistent/path/history.json")
	if err != nil {
		t.Fatalf("expected no error for missing file, got: %v", err)
	}
	if len(h.Entries) != 0 {
		t.Errorf("expected empty history, got %d entries", len(h.Entries))
	}
}

func TestDefaultPath_NotEmpty(t *testing.T) {
	p := history.DefaultPath()
	if p == "" {
		t.Error("expected non-empty default path")
	}
}

func TestAdd_SetsValidFlag(t *testing.T) {
	h := &history.History{}
	h.Add("bad", "", false)
	if h.Entries[0].Valid {
		t.Error("expected Valid=false")
	}
	_ = os.Stderr // suppress unused import
}
