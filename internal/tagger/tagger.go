// Package tagger provides functionality for tagging and categorizing cron expressions.
package tagger

import (
	"strings"

	"github.com/croncheck/internal/parser"
)

// Tag represents a label assigned to a cron expression.
type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// TaggedExpression associates a cron expression with one or more tags.
type TaggedExpression struct {
	Expression string `json:"expression"`
	Tags       []Tag  `json:"tags"`
}

// AutoTag analyzes a cron expression and returns suggested tags based on its schedule pattern.
func AutoTag(expr string) ([]Tag, error) {
	schedule, err := parser.Parse(expr)
	if err != nil {
		return nil, err
	}
	_ = schedule

	var tags []Tag
	fields := strings.Fields(expr)
	if len(fields) != 5 {
		return tags, nil
	}

	minute, hour, dom, month, dow := fields[0], fields[1], fields[2], fields[3], fields[4]

	if minute == "*" && hour == "*" && dom == "*" && month == "*" && dow == "*" {
		tags = append(tags, Tag{Name: "frequent", Description: "Runs every minute"})
	}
	if minute == "0" && hour == "*" {
		tags = append(tags, Tag{Name: "hourly", Description: "Runs at the top of every hour"})
	}
	if minute == "0" && hour == "0" {
		tags = append(tags, Tag{Name: "daily", Description: "Runs once a day at midnight"})
	}
	if minute == "0" && hour == "0" && dom == "*" && month == "*" && dow == "0" {
		tags = append(tags, Tag{Name: "weekly", Description: "Runs once a week on Sunday"})
	}
	if minute == "0" && hour == "0" && dom == "1" {
		tags = append(tags, Tag{Name: "monthly", Description: "Runs on the first day of each month"})
	}
	if strings.Contains(minute, "*/") {
		tags = append(tags, Tag{Name: "interval", Description: "Runs at a fixed minute interval"})
	}
	if dow != "*" {
		tags = append(tags, Tag{Name: "weekday-specific", Description: "Runs on specific days of the week"})
	}
	if month != "*" {
		tags = append(tags, Tag{Name: "month-specific", Description: "Runs in specific months"})
	}

	return tags, nil
}

// Apply attaches the provided tags to an expression, returning a TaggedExpression.
func Apply(expr string, tags []Tag) TaggedExpression {
	return TaggedExpression{
		Expression: expr,
		Tags:       tags,
	}
}
