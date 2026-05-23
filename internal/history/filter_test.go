package history_test

import (
	"testing"

	"github.com/user/croncheck/internal/history"
)

func makeHistory(entries ...history.Entry) *history.History {
	return &history.History{Entries: entries}
}

func entry(expr string, valid bool) history.Entry {
	return history.Entry{Expression: expr, Valid: valid}
}

func TestFilterValid_ReturnsOnlyValid(t *testing.T) {
	h := makeHistory(entry("* * * * *", true), entry("bad", false), entry("0 * * * *", true))
	result := history.FilterValid(h)
	if len(result) != 2 {
		t.Errorf("expected 2 valid entries, got %d", len(result))
	}
}

func TestFilterInvalid_ReturnsOnlyInvalid(t *testing.T) {
	h := makeHistory(entry("* * * * *", true), entry("bad", false))
	result := history.FilterInvalid(h)
	if len(result) != 1 {
		t.Errorf("expected 1 invalid entry, got %d", len(result))
	}
	if result[0].Expression != "bad" {
		t.Errorf("unexpected expression: %s", result[0].Expression)
	}
}

func TestFilterByExpression_MatchesExact(t *testing.T) {
	h := makeHistory(entry("* * * * *", true), entry("0 * * * *", true), entry("* * * * *", true))
	result := history.FilterByExpression(h, "* * * * *")
	if len(result) != 2 {
		t.Errorf("expected 2 matches, got %d", len(result))
	}
}

func TestDeduplicate_RemovesDuplicates(t *testing.T) {
	h := makeHistory(
		entry("* * * * *", true),
		entry("0 * * * *", true),
		entry("* * * * *", true),
	)
	result := history.Deduplicate(h)
	if len(result) != 2 {
		t.Errorf("expected 2 unique entries, got %d", len(result))
	}
}

func TestDeduplicate_KeepsMostRecent(t *testing.T) {
	h := makeHistory(
		entry("* * * * *", false),
		entry("* * * * *", true),
	)
	result := history.Deduplicate(h)
	if len(result) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(result))
	}
	if !result[0].Valid {
		t.Error("expected most recent (valid) entry to be kept")
	}
}

func TestFilterValid_EmptyHistory(t *testing.T) {
	h := &history.History{}
	result := history.FilterValid(h)
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d", len(result))
	}
}
