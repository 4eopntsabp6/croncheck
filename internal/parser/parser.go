package parser

import (
	"fmt"
	"strings"

	"github.com/robfig/cron/v3"
)

// CronExpression holds a parsed cron expression and its metadata.
type CronExpression struct {
	Raw      string
	Fields   []string
	Schedule cron.Schedule
}

// FieldNames maps positional index to cron field name.
var FieldNames = []string{"Minute", "Hour", "Day-of-Month", "Month", "Day-of-Week"}

// Parse validates and parses a standard 5-field cron expression.
func Parse(expr string) (*CronExpression, error) {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return nil, fmt.Errorf("cron expression must not be empty")
	}

	fields := strings.Fields(expr)
	if len(fields) != 5 {
		return nil, fmt.Errorf("expected 5 fields, got %d: %q", len(fields), expr)
	}

	parser := cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow,
	)

	schedule, err := parser.Parse(expr)
	if err != nil {
		return nil, fmt.Errorf("invalid cron expression %q: %w", expr, err)
	}

	return &CronExpression{
		Raw:      expr,
		Fields:   fields,
		Schedule: schedule,
	}, nil
}

// Describe returns a human-readable breakdown of each cron field.
func (c *CronExpression) Describe() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Expression: %s\n", c.Raw))
	for i, name := range FieldNames {
		sb.WriteString(fmt.Sprintf("  %-16s %s\n", name+":", c.Fields[i]))
	}
	return sb.String()
}
