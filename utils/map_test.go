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
