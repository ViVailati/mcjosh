package utils

// Map receives a slice and a function and returns a slice of the return type of the received function
func Map[S any, T any](slice []S, fn func(S) T) []T {
	if slice == nil {
		return nil
	}

	if len(slice) == 0 {
		return []T{}
	}

	result := make([]T, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// Filter receives a slice and a function and returns a slice of the elements that satisfy the received function
func Filter[S any](slice []S, fn func(S) bool) []S {
	if slice == nil {
		return nil
	}

	if len(slice) == 0 {
		return []S{}
	}

	result := make([]S, 0, len(slice))
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
