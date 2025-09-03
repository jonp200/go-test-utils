package utils

// AllStrings checks if a slice contains only strings and returns a converted copy.
func AllStrings(xi []any) ([]string, bool) {
	xs := make([]string, len(xi))

	for i, v := range xi {
		s, ok := v.(string)

		if !ok {
			return nil, false
		}

		xs[i] = s
	}

	return xs, true
}
