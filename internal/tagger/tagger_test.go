package tagger_test

import (
	"testing"

	"github.com/croncheck/internal/tagger"
)

func TestAutoTag_EveryMinute(t *testing.T) {
	tags, err := tagger.AutoTag("* * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !containsTag(tags, "frequent") {
		t.Errorf("expected tag 'frequent', got %v", tags)
	}
}

func TestAutoTag_Hourly(t *testing.T) {
	tags, err := tagger.AutoTag("0 * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !containsTag(tags, "hourly") {
		t.Errorf("expected tag 'hourly', got %v", tags)
	}
}

func TestAutoTag_Daily(t *testing.T) {
	tags, err := tagger.AutoTag("0 0 * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !containsTag(tags, "daily") {
		t.Errorf("expected tag 'daily', got %v", tags)
	}
}

func TestAutoTag_Monthly(t *testing.T) {
	tags, err := tagger.AutoTag("0 0 1 * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !containsTag(tags, "monthly") {
		t.Errorf("expected tag 'monthly', got %v", tags)
	}
}

func TestAutoTag_Interval(t *testing.T) {
	tags, err := tagger.AutoTag("*/15 * * * *")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !containsTag(tags, "interval") {
		t.Errorf("expected tag 'interval', got %v", tags)
	}
}

func TestAutoTag_InvalidExpression(t *testing.T) {
	_, err := tagger.AutoTag("not a cron")
	if err == nil {
		t.Error("expected error for invalid expression, got nil")
	}
}

func TestApply_StoresTags(t *testing.T) {
	tags := []tagger.Tag{{Name: "custom", Description: "a custom tag"}}
	te := tagger.Apply("0 * * * *", tags)
	if te.Expression != "0 * * * *" {
		t.Errorf("expected expression to be preserved, got %q", te.Expression)
	}
	if len(te.Tags) != 1 || te.Tags[0].Name != "custom" {
		t.Errorf("expected tags to be stored, got %v", te.Tags)
	}
}

func containsTag(tags []tagger.Tag, name string) bool {
	for _, t := range tags {
		if t.Name == name {
			return true
		}
	}
	return false
}
