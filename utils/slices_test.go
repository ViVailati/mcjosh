package utils

import "testing"

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		mapper   func(x int) int
		expected []int
	}{
		{
			name:     "empty slice",
			slice:    []int{},
			mapper:   func(x int) int { return x * 2 },
			expected: []int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			mapper:   func(x int) int { return x * 2 },
			expected: nil,
		},
		{
			name:     "mapping function returns correct value",
			slice:    []int{1, 2, 3},
			mapper:   func(x int) int { return x * 2 },
			expected: []int{2, 4, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Map(tt.slice, tt.mapper)
			if tt.expected == nil {
				if actual != nil {
					t.Errorf("expected nil, but got %v", actual)
				}
			} else {
				for i, v := range actual {
					if v != tt.expected[i] {
						t.Errorf("expected %v, but got %v", tt.expected, actual)
					}
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		filter   func(x int) bool
		expected []int
	}{
		{
			name:     "empty slice",
			slice:    []int{},
			filter:   func(x int) bool { return x%2 == 0 },
			expected: []int{},
		},
		{
			name:     "nil slice",
			slice:    nil,
			filter:   func(x int) bool { return x%2 == 0 },
			expected: nil,
		},
		{
			name:     "filtering function returns correct value",
			slice:    []int{1, 2, 3, 4, 5},
			filter:   func(x int) bool { return x%2 == 0 },
			expected: []int{2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Filter(tt.slice, tt.filter)
			if tt.expected == nil {
				if actual != nil {
					t.Errorf("expected nil, but got %v", actual)
				}
			} else {
				for i, v := range actual {
					if v != tt.expected[i] {
						t.Errorf("expected %v, but got %v", tt.expected, actual)
					}
				}
			}
		})
	}
}
