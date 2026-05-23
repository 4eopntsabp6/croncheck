package differ_test

import (
	"testing"
	"time"

	"github.com/user/croncheck/internal/differ"
)

var epoch = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

func TestCompare_SameExpression(t *testing.T) {
	d := differ.Compare("0 * * * *", "0 * * * *", epoch, 3)
	if !d.Valid {
		t.Fatalf("expected valid diff, got errors: %v", d.Errors)
	}
	for _, fd := range d.FieldDiffs {
		if !fd.Same {
			t.Errorf("field %s should be same, got A=%s B=%s", fd.Field, fd.ValueA, fd.ValueB)
		}
	}
	for _, rd := range d.NextRunDiffs {
		if rd.Delta != 0 {
			t.Errorf("run %d: expected zero delta, got %v", rd.Index, rd.Delta)
		}
	}
}

func TestCompare_DifferentExpressions(t *testing.T) {
	d := differ.Compare("0 * * * *", "30 * * * *", epoch, 2)
	if !d.Valid {
		t.Fatalf("expected valid diff")
	}
	minField := d.FieldDiffs[0]
	if minField.Same {
		t.Error("minute field should differ")
	}
	if minField.ValueA != "0" || minField.ValueB != "30" {
		t.Errorf("unexpected minute values: %s %s", minField.ValueA, minField.ValueB)
	}
	for _, rd := range d.NextRunDiffs {
		if rd.Delta == 0 {
			t.Errorf("run %d: expected non-zero delta", rd.Index)
		}
	}
}

func TestCompare_InvalidExpressionA(t *testing.T) {
	d := differ.Compare("invalid", "* * * * *", epoch, 1)
	if d.Valid {
		t.Error("expected invalid diff")
	}
	if len(d.Errors) == 0 {
		t.Error("expected at least one error")
	}
}

func TestCompare_BothInvalid(t *testing.T) {
	d := differ.Compare("bad", "also bad", epoch, 1)
	if d.Valid {
		t.Error("expected invalid diff")
	}
	if len(d.Errors) < 2 {
		t.Errorf("expected 2 errors, got %d", len(d.Errors))
	}
}

func TestCompare_RunCount(t *testing.T) {
	d := differ.Compare("* * * * *", "* * * * *", epoch, 5)
	if len(d.NextRunDiffs) != 5 {
		t.Errorf("expected 5 run diffs, got %d", len(d.NextRunDiffs))
	}
}
