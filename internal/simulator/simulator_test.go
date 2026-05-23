package simulator_test

import (
	"testing"
	"time"

	"github.com/croncheck/internal/parser"
	"github.com/croncheck/internal/simulator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mustParse(t *testing.T, expr string) *parser.CronExpression {
	t.Helper()
	ce, err := parser.Parse(expr)
	require.NoError(t, err)
	return ce
}

func TestNextRuns_Count(t *testing.T) {
	ce := mustParse(t, "* * * * *")
	from := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	results := simulator.NextRuns(ce, from, 5)
	assert.Len(t, results, 5)
	for i, r := range results {
		assert.Equal(t, i+1, r.N)
	}
}

func TestNextRuns_Ordering(t *testing.T) {
	ce := mustParse(t, "0 * * * *")
	from := time.Date(2024, 6, 15, 10, 30, 0, 0, time.UTC)
	results := simulator.NextRuns(ce, from, 3)
	require.Len(t, results, 3)
	assert.True(t, results[0].Time.Before(results[1].Time))
	assert.True(t, results[1].Time.Before(results[2].Time))
	// First run should be at 11:00
	assert.Equal(t, 11, results[0].Time.Hour())
	assert.Equal(t, 0, results[0].Time.Minute())
}

func TestNextRuns_Zero(t *testing.T) {
	ce := mustParse(t, "* * * * *")
	from := time.Now()
	results := simulator.NextRuns(ce, from, 0)
	assert.Empty(t, results)
}

func TestFormat_Output(t *testing.T) {
	ce := mustParse(t, "0 12 * * *")
	from := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	results := simulator.NextRuns(ce, from, 2)
	output := simulator.Format(results, time.UTC)
	assert.Contains(t, output, "#")
	assert.Contains(t, output, "Scheduled Time")
	assert.Contains(t, output, "2024-03-01 12:00:00 UTC")
	assert.Contains(t, output, "2024-03-02 12:00:00 UTC")
}
