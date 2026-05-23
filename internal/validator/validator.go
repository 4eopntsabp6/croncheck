package validator

import (
	"fmt"
	"strconv"
	"strings"
)

// ValidationError holds a field name and a human-readable reason.
type ValidationError struct {
	Field  string
	Reason string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("invalid %s field: %s", e.Field, e.Reason)
}

// Result aggregates all errors found in a single expression.
type Result struct {
	Valid  bool
	Errors []ValidationError
}

var fieldMeta = []struct {
	name string
	min  int
	max  int
}{
	{"minute", 0, 59},
	{"hour", 0, 23},
	{"day-of-month", 1, 31},
	{"month", 1, 12},
	{"day-of-week", 0, 6},
}

// Validate checks each field of a standard 5-part cron expression.
func Validate(expr string) Result {
	parts := strings.Fields(expr)
	if len(parts) != 5 {
		return Result{
			Valid:  false,
			Errors: []ValidationError{{Field: "expression", Reason: fmt.Sprintf("expected 5 fields, got %d", len(parts))}},
		}
	}

	var errs []ValidationError
	for i, part := range parts {
		meta := fieldMeta[i]
		if err := validateField(part, meta.name, meta.min, meta.max); err != nil {
			errs = append(errs, *err)
		}
	}

	return Result{Valid: len(errs) == 0, Errors: errs}
}

func validateField(field, name string, min, max int) *ValidationError {
	if field == "*" {
		return nil
	}
	// Handle step values: */n or range/n
	if strings.Contains(field, "/") {
		parts := strings.SplitN(field, "/", 2)
		step, err := strconv.Atoi(parts[1])
		if err != nil || step < 1 {
			return &ValidationError{Field: name, Reason: fmt.Sprintf("invalid step value %q", parts[1])}
		}
		if parts[0] != "*" {
			return validateField(parts[0], name, min, max)
		}
		return nil
	}
	// Handle lists: 1,2,3
	for _, item := range strings.Split(field, ",") {
		// Handle ranges: 1-5
		if strings.Contains(item, "-") {
			bounds := strings.SplitN(item, "-", 2)
			lo, err1 := strconv.Atoi(bounds[0])
			hi, err2 := strconv.Atoi(bounds[1])
			if err1 != nil || err2 != nil || lo > hi || lo < min || hi > max {
				return &ValidationError{Field: name, Reason: fmt.Sprintf("invalid range %q (allowed %d-%d)", item, min, max)}
			}
			continue
		}
		v, err := strconv.Atoi(item)
		if err != nil || v < min || v > max {
			return &ValidationError{Field: name, Reason: fmt.Sprintf("value %q out of range %d-%d", item, min, max)}
		}
	}
	return nil
}
