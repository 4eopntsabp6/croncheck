package describer

import (
	"strings"
	"testing"
)

func TestDescribe_EveryMinute(t *testing.T) {
	desc, err := Describe("* * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(desc, "every minute") {
		t.Errorf("expected 'every minute' in %q", desc)
	}
}

func TestDescribe_HourlyStep(t *testing.T) {
	desc, err := Describe("0 */2 * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(desc, "every 2 hour") {
		t.Errorf("expected step description in %q", desc)
	}
}

func TestDescribe_SpecificMonth(t *testing.T) {
	desc, err := Describe("0 9 1 3 *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(desc, "March") {
		t.Errorf("expected 'March' in %q", desc)
	}
}

func TestDescribe_SpecificWeekday(t *testing.T) {
	desc, err := Describe("0 9 * * 1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(desc, "Monday") {
		t.Errorf("expected 'Monday' in %q", desc)
	}
}

func TestDescribe_InvalidExpression(t *testing.T) {
	_, err := Describe("not a cron")
	if err == nil {
		t.Fatal("expected error for invalid expression")
	}
}

func TestDescribe_WrongFieldCount(t *testing.T) {
	_, err := Describe("* * *")
	if err == nil {
		t.Fatal("expected error for wrong field count")
	}
}

func TestDescribe_RangeField(t *testing.T) {
	desc, err := Describe("0 9-17 * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(desc, "9") || !strings.Contains(desc, "17") {
		t.Errorf("expected range values in %q", desc)
	}
}
