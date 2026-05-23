// Package formatter provides human-readable formatting for cron expression summaries.
package formatter

import (
	"fmt"
	"strings"
	"time"

	"github.com/user/croncheck/internal/describer"
	"github.com/user/croncheck/internal/simulator"
	"github.com/user/croncheck/internal/validator"

	"github.com/robfig/cron/v3"
)

// Summary holds a structured summary of a cron expression.
type Summary struct {
	Expression  string
	Valid       bool
	Errors      []string
	Description string
	NextRuns    []time.Time
}

// Summarize parses and describes a cron expression, returning a Summary.
func Summarize(expr string, count int, from time.Time) Summary {
	errs := validator.Validate(expr)
	if len(errs) > 0 {
		messages := make([]string, len(errs))
		for i, e := range errs {
			messages[i] = e.Error()
		}
		return Summary{
			Expression: expr,
			Valid:      false,
			Errors:     messages,
		}
	}

	schedule, err := cron.ParseStandard(expr)
	if err != nil {
		return Summary{
			Expression: expr,
			Valid:      false,
			Errors:     []string{err.Error()},
		}
	}

	desc := describer.Describe(expr)
	runs := simulator.NextRuns(schedule, from, count)

	return Summary{
		Expression:  expr,
		Valid:       true,
		Errors:      nil,
		Description: desc,
		NextRuns:    runs,
	}
}

// FormatSummary renders a Summary as a human-readable multi-line string.
func FormatSummary(s Summary) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Expression : %s\n", s.Expression))

	if !s.Valid {
		sb.WriteString("Valid      : false\n")
		sb.WriteString("Errors     :\n")
		for _, e := range s.Errors {
			sb.WriteString(fmt.Sprintf("  - %s\n", e))
		}
		return sb.String()
	}

	sb.WriteString("Valid      : true\n")
	sb.WriteString(fmt.Sprintf("Description: %s\n", s.Description))
	sb.WriteString("Next Runs  :\n")
	for i, r := range s.NextRuns {
		sb.WriteString(fmt.Sprintf("  %2d. %s\n", i+1, simulator.Format(r)))
	}

	return sb.String()
}
