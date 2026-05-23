package exporter

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"
)

var baseTime = time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)

func sampleRuns() []time.Time {
	return []time.Time{
		baseTime.Add(1 * time.Hour),
		baseTime.Add(2 * time.Hour),
		baseTime.Add(25 * time.Hour),
	}
}

func TestExportJSON_ValidOutput(t *testing.T) {
	var buf bytes.Buffer
	err := ExportJSON(&buf, sampleRuns(), baseTime)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var entries []RunEntry
	if err := json.Unmarshal(buf.Bytes(), &entries); err != nil {
		t.Fatalf("failed to parse JSON output: %v", err)
	}
	if len(entries) != 3 {
		t.Errorf("expected 3 entries, got %d", len(entries))
	}
	if entries[0].Index != 1 {
		t.Errorf("expected first index 1, got %d", entries[0].Index)
	}
	if !strings.Contains(entries[0].Relative, "in") {
		t.Errorf("expected relative to contain 'in', got %q", entries[0].Relative)
	}
}

func TestExportCSV_ValidOutput(t *testing.T) {
	var buf bytes.Buffer
	err := ExportCSV(&buf, sampleRuns(), baseTime)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	// header + 3 data rows
	if len(lines) != 4 {
		t.Errorf("expected 4 lines, got %d", len(lines))
	}
	if !strings.HasPrefix(lines[0], "index") {
		t.Errorf("expected header line, got %q", lines[0])
	}
	if !strings.Contains(lines[1], "1") {
		t.Errorf("expected first data row to contain index 1")
	}
}

func TestExportCSV_EmptyRuns(t *testing.T) {
	var buf bytes.Buffer
	err := ExportCSV(&buf, []time.Time{}, baseTime)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	if len(lines) != 1 {
		t.Errorf("expected only header line, got %d lines", len(lines))
	}
}

func TestFormatRelative_Future(t *testing.T) {
	now := baseTime
	future := baseTime.Add(90 * time.Minute)
	result := formatRelative(future, now)
	if !strings.HasPrefix(result, "in") {
		t.Errorf("expected 'in ...', got %q", result)
	}
}

func TestFormatRelative_Past(t *testing.T) {
	now := baseTime
	past := baseTime.Add(-30 * time.Minute)
	result := formatRelative(past, now)
	if !strings.HasSuffix(result, "ago") {
		t.Errorf("expected '... ago', got %q", result)
	}
}
