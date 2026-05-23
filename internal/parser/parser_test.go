package parser_test

import (
	"testing"

	"github.com/croncheck/internal/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse_Valid(t *testing.T) {
	tests := []struct {
		name   string
		expr   string
		wantFields []string
	}{
		{"every minute", "* * * * *", []string{"*", "*", "*", "*", "*"}},
		{"every hour", "0 * * * *", []string{"0", "*", "*", "*", "*"}},
		{"daily midnight", "0 0 * * *", []string{"0", "0", "*", "*", "*"}},
		{"weekdays 9am", "0 9 * * 1-5", []string{"0", "9", "*", "*", "1-5"}},
		{"step expression", "*/15 * * * *", []string{"*/15", "*", "*", "*", "*"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ce, err := parser.Parse(tt.expr)
			require.NoError(t, err)
			assert.Equal(t, tt.expr, ce.Raw)
			assert.Equal(t, tt.wantFields, ce.Fields)
			assert.NotNil(t, ce.Schedule)
		})
	}
}

func TestParse_Invalid(t *testing.T) {
	tests := []struct {
		name string
		expr string
	}{
		{"empty string", ""},
		{"too few fields", "* * * *"},
		{"too many fields", "* * * * * *"},
		{"bad minute range", "60 * * * *"},
		{"bad hour value", "0 25 * * *"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := parser.Parse(tt.expr)
			assert.Error(t, err)
		})
	}
}

func TestDescribe(t *testing.T) {
	ce, err := parser.Parse("0 9 * * 1-5")
	require.NoError(t, err)
	desc := ce.Describe()
	assert.Contains(t, desc, "0 9 * * 1-5")
	assert.Contains(t, desc, "Minute")
	assert.Contains(t, desc, "Day-of-Week")
}
