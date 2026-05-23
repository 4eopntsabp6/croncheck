// Package exporter provides functionality to export cron schedule data
// to various formats such as JSON and CSV.
package exporter

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// RunEntry represents a single scheduled run entry for export.
type RunEntry struct {
	Index    int    `json:"index"`
	Time     string `json:"time"`
	Relative string `json:"relative"`
}

// ExportJSON writes the scheduled runs as a JSON array to the given writer.
func ExportJSON(w io.Writer, runs []time.Time, now time.Time) error {
	entries := buildEntries(runs, now)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(entries); err != nil {
		return fmt.Errorf("exporter: json encode failed: %w", err)
	}
	return nil
}

// ExportCSV writes the scheduled runs as CSV rows to the given writer.
// The CSV includes a header row: index, time, relative.
func ExportCSV(w io.Writer, runs []time.Time, now time.Time) error {
	entries := buildEntries(runs, now)
	cw := csv.NewWriter(w)
	if err := cw.Write([]string{"index", "time", "relative"}); err != nil {
		return fmt.Errorf("exporter: csv write header failed: %w", err)
	}
	for _, e := range entries {
		row := []string{
			fmt.Sprintf("%d", e.Index),
			e.Time,
			e.Relative,
		}
		if err := cw.Write(row); err != nil {
			return fmt.Errorf("exporter: csv write row failed: %w", err)
		}
	}
	cw.Flush()
	return cw.Error()
}

// buildEntries converts a slice of times into RunEntry values.
func buildEntries(runs []time.Time, now time.Time) []RunEntry {
	entries := make([]RunEntry, 0, len(runs))
	for i, t := range runs {
		entries = append(entries, RunEntry{
			Index:    i + 1,
			Time:     t.Format("2006-01-02 15:04:05 MST"),
			Relative: formatRelative(t, now),
		})
	}
	return entries
}

// formatRelative returns a human-readable relative time string.
func formatRelative(t, now time.Time) string {
	d := t.Sub(now)
	if d < 0 {
		d = -d
		return fmt.Sprintf("%s ago", formatDuration(d))
	}
	return fmt.Sprintf("in %s", formatDuration(d))
}

func formatDuration(d time.Duration) string {
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	if d < time.Hour {
		return fmt.Sprintf("%dm", int(d.Minutes()))
	}
	if d < 24*time.Hour {
		return fmt.Sprintf("%dh%dm", int(d.Hours()), int(d.Minutes())%60)
	}
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	return fmt.Sprintf("%dd%dh", days, hours)
}
