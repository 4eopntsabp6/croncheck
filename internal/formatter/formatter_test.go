package formatter_test

import (
	"strings"
	"testing"
	"time"

	"github.com/user/croncheck/internal/formatter"
)

var fixedNow = time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)

func TestSummarize_ValidExpression(t *testing.T) {
	s := formatter.Summarize("0 * * * *", 3, fixedNow)

	if !s.Valid {
		t.Fatalf("expected valid summary, got errors: %v", s.Errors)
	}
	if s.Expression != "0 * * * *" {
		t.Errorf("unexpected expression: %s", s.Expression)
	}
	if s.Description == "" {
		t.Error("expected non-empty description")
	}
	if len(s.NextRuns) != 3 {
		t.Errorf("expected 3 next runs, got %d", len(s.NextRuns))
	}
	if len(s.Errors) != 0 {
		t.Errorf("expected no errors, got %v", s.Errors)
	}
}

func TestSummarize_InvalidExpression(t *testing.T) {
	s := formatter.Summarize("invalid expr", 5, fixedNow)

	if s.Valid {
		t.Fatal("expected invalid summary")
	}
	if len(s.Errors) == 0 {
		t.Error("expected at least one error")
	}
	if len(s.NextRuns) != 0 {
		t.Errorf("expected no next runs for invalid expr, got %d", len(s.NextRuns))
	}
}

func TestFormatSummary_ValidContainsFields(t *testing.T) {
	s := formatter.Summarize("*/5 * * * *", 2, fixedNow)
	out := formatter.FormatSummary(s)

	for _, want := range []string{"Expression", "Valid", "Description", "Next Runs"} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing field %q", want)
		}
	}
}

func TestFormatSummary_InvalidShowsErrors(t *testing.T) {
	s := formatter.Summarize("bad cron", 3, fixedNow)
	out := formatter.FormatSummary(s)

	if !strings.Contains(out, "false") {
		t.Error("expected 'false' in output for invalid expression")
	}
	if !strings.Contains(out, "Errors") {
		t.Error("expected 'Errors' section in output")
	}
}

func TestFormatSummary_NextRunCount(t *testing.T) {
	s := formatter.Summarize("0 9 * * 1", 4, fixedNow)
	out := formatter.FormatSummary(s)

	count := strings.Count(out, ".")
	if count < 4 {
		t.Errorf("expected at least 4 run entries in output, found %d dots", count)
	}
}
