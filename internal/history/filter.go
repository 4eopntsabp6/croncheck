package history

// FilterValid returns only entries where Valid is true.
func FilterValid(h *History) []Entry {
	return filterBy(h, func(e Entry) bool { return e.Valid })
}

// FilterInvalid returns only entries where Valid is false.
func FilterInvalid(h *History) []Entry {
	return filterBy(h, func(e Entry) bool { return !e.Valid })
}

// FilterByExpression returns entries whose expression matches the given string.
func FilterByExpression(h *History, expr string) []Entry {
	return filterBy(h, func(e Entry) bool { return e.Expression == expr })
}

// Deduplicate returns entries with duplicate expressions removed,
// keeping the most recent occurrence of each expression.
func Deduplicate(h *History) []Entry {
	seen := make(map[string]bool)
	result := []Entry{}
	for i := len(h.Entries) - 1; i >= 0; i-- {
		e := h.Entries[i]
		if !seen[e.Expression] {
			seen[e.Expression] = true
			result = append([]Entry{e}, result...)
		}
	}
	return result
}

func filterBy(h *History, predicate func(Entry) bool) []Entry {
	result := []Entry{}
	for _, e := range h.Entries {
		if predicate(e) {
			result = append(result, e)
		}
	}
	return result
}
