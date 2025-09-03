package utils

import (
	"encoding/json"
	"reflect"
	"sort"
)

// NormalizeJSON sorts arrays in JSON to ensure order-independent comparison.
func NormalizeJSON(i any) any {
	switch obj := i.(type) {
	case map[string]any:
		// Recursively normalize
		for k, v := range obj {
			obj[k] = NormalizeJSON(v)
		}

		return obj
	case []any:
		sorted := make([]any, len(obj))

		copy(sorted, obj)

		// If the slice contains only strings, use `sort.Strings` to sort them
		if strSlice, ok := AllStrings(sorted); ok {
			sort.Strings(strSlice)

			for ix, s := range strSlice {
				sorted[ix] = s
			}
		} else {
			for ix, v := range sorted {
				sorted[ix] = NormalizeJSON(v)
			}
		}

		return sorted
	default:
		return i
	}
}

// CompareJSONBytes checks if two JSON byte slices are equal regardless of field order and array item order.
func CompareJSONBytes(a, b []byte) (bool, error) {
	var objA, objB any

	if err := json.Unmarshal(a, &objA); err != nil {
		return false, err
	}

	if err := json.Unmarshal(b, &objB); err != nil {
		return false, err
	}

	// Normalize (sort arrays inside)
	objA = NormalizeJSON(objA)
	objB = NormalizeJSON(objB)

	return reflect.DeepEqual(objA, objB), nil
}
