package visualizer

import (
	"strings"
	"testing"
	"time"
)

func makeRuns(base time.Time, intervals ...time.Duration) []time.Time {
	runs := make([]time.Time, len(intervals))
	for i, d := range intervals {
		runs[i] = base.Add(d)
	}
	return runs
}

func TestTimeline_ContainsMarkers(t *testing.T) {
	base := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	runs := makeRuns(base, 0, 15*time.Minute, 30*time.Minute, 45*time.Minute)
	out := Timeline(runs, time.Hour)
	if !strings.Contains(out, "*") {
		t.Errorf("expected '*' markers in timeline, got:\n%s", out)
	}
}

func TestTimeline_EmptyRuns(t *testing.T) {
	out := Timeline(nil, time.Hour)
	if !strings.Contains(out, "no runs") {
		t.Errorf("expected 'no runs' message, got: %q", out)
	}
}

func TestTimeline_HasBorders(t *testing.T) {
	base := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	runs := makeRuns(base, 10*time.Minute)
	out := Timeline(runs, time.Hour)
	if !strings.Contains(out, "|") {
		t.Errorf("expected '|' borders in timeline, got:\n%s", out)
	}
}

func TestTable_ContainsRunNumbers(t *testing.T) {
	base := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	runs := makeRuns(base, 5*time.Minute, 10*time.Minute, 15*time.Minute)
	out := Table(runs, base)
	for _, num := range []string{"1", "2", "3"} {
		if !strings.Contains(out, num) {
			t.Errorf("expected run number %s in table output", num)
		}
	}
}

func TestTable_EmptyRuns(t *testing.T) {
	out := Table(nil, time.Now())
	if !strings.Contains(out, "No scheduled") {
		t.Errorf("expected empty message, got: %q", out)
	}
}

func TestFormatDuration_Hours(t *testing.T) {
	d := 2*time.Hour + 30*time.Minute
	result := formatDuration(d)
	if !strings.Contains(result, "2h") {
		t.Errorf("expected hours in result, got: %q", result)
	}
}

func TestFormatDuration_Past(t *testing.T) {
	result := formatDuration(-1 * time.Minute)
	if result != "in the past" {
		t.Errorf("expected 'in the past', got: %q", result)
	}
}
