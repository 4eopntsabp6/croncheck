package visualizer

import (
	"fmt"
	"strings"
	"time"
)

const (
	// TimelineWidth is the number of characters used for the timeline bar.
	TimelineWidth = 60
)

// Timeline renders an ASCII timeline of upcoming run times within a window.
func Timeline(runs []time.Time, window time.Duration) string {
	if len(runs) == 0 {
		return "(no runs scheduled)"
	}

	now := runs[0]
	end := now.Add(window)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Timeline: %s → %s\n", now.Format("15:04"), end.Format("15:04")))
	sb.WriteString("|" + strings.Repeat("-", TimelineWidth) + "|\n")

	bar := []rune(strings.Repeat(" ", TimelineWidth))
	for _, r := range runs {
		offset := r.Sub(now)
		if offset < 0 || offset > window {
			continue
		}
		pos := int(float64(offset) / float64(window) * float64(TimelineWidth))
		if pos >= TimelineWidth {
			pos = TimelineWidth - 1
		}
		bar[pos] = '*'
	}

	sb.WriteString("|" + string(bar) + "|\n")
	return sb.String()
}

// Table renders a formatted table of run times with their relative offsets.
func Table(runs []time.Time, from time.Time) string {
	if len(runs) == 0 {
		return "No scheduled runs found.\n"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%-5s  %-25s  %s\n", "#", "Scheduled Time", "In"))
	sb.WriteString(strings.Repeat("-", 50) + "\n")

	for i, r := range runs {
		diff := r.Sub(from).Round(time.Second)
		sb.WriteString(fmt.Sprintf("%-5d  %-25s  %s\n",
			i+1,
			r.Format("2006-01-02 15:04:05 MST"),
			formatDuration(diff),
		))
	}
	return sb.String()
}

func formatDuration(d time.Duration) string {
	if d < 0 {
		return "in the past"
	}
	h := int(d.Hours())
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	if h > 0 {
		return fmt.Sprintf("%dh %dm", h, m)
	}
	if m > 0 {
		return fmt.Sprintf("%dm %ds", m, s)
	}
	return fmt.Sprintf("%ds", s)
}
