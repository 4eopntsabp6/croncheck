package simulator

import (
	"fmt"
	"strings"
	"time"

	"github.com/croncheck/internal/parser"
)

// RunResult holds a single simulated next-run entry.
type RunResult struct {
	N    int
	Time time.Time
}

// NextRuns computes the next n scheduled times starting from `from`.
func NextRuns(ce *parser.CronExpression, from time.Time, n int) []RunResult {
	if n <= 0 {
		return nil
	}
	results := make([]RunResult, 0, n)
	t := from
	for i := 1; i <= n; i++ {
		t = ce.Schedule.Next(t)
		results = append(results, RunResult{N: i, Time: t})
	}
	return results
}

// Format renders a slice of RunResults as a formatted table string.
func Format(results []RunResult, loc *time.Location) string {
	if loc == nil {
		loc = time.UTC
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%-4s  %-30s  %s\n", "#", "Scheduled Time", "Relative"))
	sb.WriteString(strings.Repeat("-", 60) + "\n")
	now := time.Now().In(loc)
	for _, r := range results {
		t := r.Time.In(loc)
		rel := formatRelative(t.Sub(now))
		sb.WriteString(fmt.Sprintf("%-4d  %-30s  %s\n", r.N, t.Format("2006-01-02 15:04:05 MST"), rel))
	}
	return sb.String()
}

func formatRelative(d time.Duration) string {
	if d < 0 {
		return "in the past"
	}
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	if h > 0 {
		return fmt.Sprintf("in %dh %dm", h, m)
	}
	if m > 0 {
		return fmt.Sprintf("in %dm %ds", m, s)
	}
	return fmt.Sprintf("in %ds", s)
}
