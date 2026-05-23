// Package differ compares two cron expressions and highlights their differences.
package differ

import (
	"fmt"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

// Diff holds the comparison result between two cron expressions.
type Diff struct {
	ExprA       string
	ExprB       string
	Valid        bool
	Errors       []string
	FieldDiffs   []FieldDiff
	NextRunDiffs []RunDiff
}

// FieldDiff describes a difference in a single cron field.
type FieldDiff struct {
	Field  string
	ValueA string
	ValueB string
	Same   bool
}

// RunDiff compares a next-run time from each expression.
type RunDiff struct {
	Index int
	RunA  time.Time
	RunB  time.Time
	Delta time.Duration
}

var fieldNames = []string{"Minute", "Hour", "Day", "Month", "Weekday"}

// Compare parses and compares two cron expressions, returning a Diff.
func Compare(exprA, exprB string, from time.Time, count int) Diff {
	d := Diff{ExprA: exprA, ExprB: exprB, Valid: true}

	parserOpts := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)

	_, errA := parserOpts.Parse(exprA)
	_, errB := parserOpts.Parse(exprB)

	if errA != nil {
		d.Valid = false
		d.Errors = append(d.Errors, fmt.Sprintf("expression A: %v", errA))
	}
	if errB != nil {
		d.Valid = false
		d.Errors = append(d.Errors, fmt.Sprintf("expression B: %v", errB))
	}
	if !d.Valid {
		return d
	}

	partsA := strings.Fields(exprA)
	partsB := strings.Fields(exprB)
	for i, name := range fieldNames {
		va, vb := "", ""
		if i < len(partsA) {
			va = partsA[i]
		}
		if i < len(partsB) {
			vb = partsB[i]
		}
		d.FieldDiffs = append(d.FieldDiffs, FieldDiff{
			Field:  name,
			ValueA: va,
			ValueB: vb,
			Same:   va == vb,
		})
	}

	schedA, _ := parserOpts.Parse(exprA)
	schedB, _ := parserOpts.Parse(exprB)

	tA, tB := from, from
	for i := 0; i < count; i++ {
		tA = schedA.Next(tA)
		tB = schedB.Next(tB)
		d.NextRunDiffs = append(d.NextRunDiffs, RunDiff{
			Index: i + 1,
			RunA:  tA,
			RunB:  tB,
			Delta: tB.Sub(tA),
		})
	}

	return d
}
