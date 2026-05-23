package validator

import (
	"testing"
)

func TestValidate_EveryMinute(t *testing.T) {
	r := Validate("* * * * *")
	if !r.Valid {
		t.Fatalf("expected valid, got errors: %v", r.Errors)
	}
}

func TestValidate_SpecificValues(t *testing.T) {
	r := Validate("30 9 15 6 1")
	if !r.Valid {
		t.Fatalf("expected valid, got errors: %v", r.Errors)
	}
}

func TestValidate_StepExpression(t *testing.T) {
	r := Validate("*/15 */2 * * *")
	if !r.Valid {
		t.Fatalf("expected valid, got errors: %v", r.Errors)
	}
}

func TestValidate_RangeExpression(t *testing.T) {
	r := Validate("0 9-17 * * 1-5")
	if !r.Valid {
		t.Fatalf("expected valid, got errors: %v", r.Errors)
	}
}

func TestValidate_ListExpression(t *testing.T) {
	r := Validate("0 8,12,18 * * *")
	if !r.Valid {
		t.Fatalf("expected valid, got errors: %v", r.Errors)
	}
}

func TestValidate_WrongFieldCount(t *testing.T) {
	r := Validate("* * * *")
	if r.Valid {
		t.Fatal("expected invalid for 4-field expression")
	}
	if len(r.Errors) == 0 {
		t.Fatal("expected at least one error")
	}
}

func TestValidate_OutOfRangeMinute(t *testing.T) {
	r := Validate("60 * * * *")
	if r.Valid {
		t.Fatal("expected invalid for minute=60")
	}
	if r.Errors[0].Field != "minute" {
		t.Fatalf("expected minute field error, got %q", r.Errors[0].Field)
	}
}

func TestValidate_InvalidStep(t *testing.T) {
	r := Validate("*/0 * * * *")
	if r.Valid {
		t.Fatal("expected invalid for step=0")
	}
}

func TestValidate_InvalidRange(t *testing.T) {
	r := Validate("* * * * 5-2")
	if r.Valid {
		t.Fatal("expected invalid for reversed range 5-2")
	}
}

func TestValidate_MultipleErrors(t *testing.T) {
	r := Validate("99 25 * * *")
	if r.Valid {
		t.Fatal("expected invalid")
	}
	if len(r.Errors) < 2 {
		t.Fatalf("expected at least 2 errors, got %d", len(r.Errors))
	}
}
