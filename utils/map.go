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
