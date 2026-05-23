package describer

import (
	"fmt"
	"strings"

	"github.com/robfig/cron/v3"
)

// Describe returns a human-readable description of a cron expression.
func Describe(expr string) (string, error) {
	_, err := cron.ParseStandard(expr)
	if err != nil {
		return "", fmt.Errorf("invalid cron expression: %w", err)
	}

	parts := strings.Fields(expr)
	if len(parts) != 5 {
		return "", fmt.Errorf("expected 5 fields, got %d", len(parts))
	}

	minute := describePart(parts[0], "minute", 0, 59)
	hour := describePart(parts[1], "hour", 0, 23)
	dom := describePart(parts[2], "day of month", 1, 31)
	month := describeMonth(parts[3])
	dow := describeWeekday(parts[4])

	return fmt.Sprintf("At %s past %s, on %s, in %s, on %s", minute, hour, dom, month, dow), nil
}

func describePart(field, unit string, min, max int) string {
	if field == "*" {
		return fmt.Sprintf("every %s", unit)
	}
	if strings.HasPrefix(field, "*/") {
		step := strings.TrimPrefix(field, "*/")
		return fmt.Sprintf("every %s %s(s)", step, unit)
	}
	if strings.Contains(field, "-") {
		parts := strings.SplitN(field, "-", 2)
		return fmt.Sprintf("%s %s through %s", unit, parts[0], parts[1])
	}
	if strings.Contains(field, ",") {
		return fmt.Sprintf("%s %s", unit, field)
	}
	return fmt.Sprintf("%s %s", unit, field)
}

func describeMonth(field string) string {
	months := map[string]string{
		"1": "January", "2": "February", "3": "March", "4": "April",
		"5": "May", "6": "June", "7": "July", "8": "August",
		"9": "September", "10": "October", "11": "November", "12": "December",
	}
	if field == "*" {
		return "every month"
	}
	if name, ok := months[field]; ok {
		return name
	}
	return field
}

func describeWeekday(field string) string {
	days := map[string]string{
		"0": "Sunday", "1": "Monday", "2": "Tuesday", "3": "Wednesday",
		"4": "Thursday", "5": "Friday", "6": "Saturday", "7": "Sunday",
	}
	if field == "*" {
		return "every day of the week"
	}
	if name, ok := days[field]; ok {
		return name
	}
	return field
}
