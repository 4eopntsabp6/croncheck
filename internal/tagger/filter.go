package tagger

// FilterByTag returns only those TaggedExpressions that contain the given tag name.
func FilterByTag(expressions []TaggedExpression, tagName string) []TaggedExpression {
	var result []TaggedExpression
	for _, te := range expressions {
		if hasTag(te, tagName) {
			result = append(result, te)
		}
	}
	return result
}

// FilterByAnyTag returns TaggedExpressions that contain at least one of the given tag names.
func FilterByAnyTag(expressions []TaggedExpression, tagNames ...string) []TaggedExpression {
	set := make(map[string]struct{}, len(tagNames))
	for _, n := range tagNames {
		set[n] = struct{}{}
	}
	var result []TaggedExpression
	for _, te := range expressions {
		for _, t := range te.Tags {
			if _, ok := set[t.Name]; ok {
				result = append(result, te)
				break
			}
		}
	}
	return result
}

// TagNames returns a deduplicated list of all tag names present across the given expressions.
func TagNames(expressions []TaggedExpression) []string {
	seen := make(map[string]struct{})
	var names []string
	for _, te := range expressions {
		for _, t := range te.Tags {
			if _, ok := seen[t.Name]; !ok {
				seen[t.Name] = struct{}{}
				names = append(names, t.Name)
			}
		}
	}
	return names
}

func hasTag(te TaggedExpression, name string) bool {
	for _, t := range te.Tags {
		if t.Name == name {
			return true
		}
	}
	return false
}
