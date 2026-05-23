package tagger_test

import (
	"testing"

	"github.com/croncheck/internal/tagger"
)

func makeTagged(expr string, tagNames ...string) tagger.TaggedExpression {
	var tags []tagger.Tag
	for _, n := range tagNames {
		tags = append(tags, tagger.Tag{Name: n})
	}
	return tagger.Apply(expr, tags)
}

func TestFilterByTag_ReturnsMatching(t *testing.T) {
	exprs := []tagger.TaggedExpression{
		makeTagged("* * * * *", "frequent"),
		makeTagged("0 * * * *", "hourly"),
		makeTagged("0 0 * * *", "daily"),
	}
	result := tagger.FilterByTag(exprs, "hourly")
	if len(result) != 1 || result[0].Expression != "0 * * * *" {
		t.Errorf("expected 1 hourly expression, got %v", result)
	}
}

func TestFilterByTag_NoMatch(t *testing.T) {
	exprs := []tagger.TaggedExpression{
		makeTagged("* * * * *", "frequent"),
	}
	result := tagger.FilterByTag(exprs, "monthly")
	if len(result) != 0 {
		t.Errorf("expected no results, got %v", result)
	}
}

func TestFilterByAnyTag_ReturnsMultiple(t *testing.T) {
	exprs := []tagger.TaggedExpression{
		makeTagged("* * * * *", "frequent"),
		makeTagged("0 * * * *", "hourly"),
		makeTagged("0 0 * * *", "daily"),
	}
	result := tagger.FilterByAnyTag(exprs, "frequent", "daily")
	if len(result) != 2 {
		t.Errorf("expected 2 results, got %d", len(result))
	}
}

func TestTagNames_Deduplicated(t *testing.T) {
	exprs := []tagger.TaggedExpression{
		makeTagged("* * * * *", "frequent", "interval"),
		makeTagged("*/5 * * * *", "interval"),
	}
	names := tagger.TagNames(exprs)
	if len(names) != 2 {
		t.Errorf("expected 2 unique tag names, got %d: %v", len(names), names)
	}
}

func TestTagNames_Empty(t *testing.T) {
	names := tagger.TagNames(nil)
	if len(names) != 0 {
		t.Errorf("expected empty result for nil input, got %v", names)
	}
}
