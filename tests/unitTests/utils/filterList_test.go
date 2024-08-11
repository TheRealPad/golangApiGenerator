package utils

import (
	"httpServer/src/utils"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	// Test cases
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
		filter   interface{}
	}{
		{
			name:     "Filter even numbers",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 6},
			filter:   func(n int) bool { return n%2 == 0 },
		},
		{
			name:     "Filter strings with more than 3 characters",
			input:    []string{"go", "golang", "java", "js", "python"},
			expected: []string{"golang", "java", "python"},
			filter:   func(s string) bool { return len(s) > 3 },
		},
		{
			name:     "Filter positive floats",
			input:    []float64{-1.5, 2.3, -3.4, 4.5, 0.0},
			expected: []float64{2.3, 4.5},
			filter:   func(f float64) bool { return f > 0 },
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.input.(type) {
			case []int:
				filtered := utils.Filter(v, tt.filter.(func(int) bool))
				if !reflect.DeepEqual(filtered, tt.expected) {
					t.Errorf("Expected %v, got %v", tt.expected, filtered)
				}
			case []string:
				filtered := utils.Filter(v, tt.filter.(func(string) bool))
				if !reflect.DeepEqual(filtered, tt.expected) {
					t.Errorf("Expected %v, got %v", tt.expected, filtered)
				}
			case []float64:
				filtered := utils.Filter(v, tt.filter.(func(float64) bool))
				if !reflect.DeepEqual(filtered, tt.expected) {
					t.Errorf("Expected %v, got %v", tt.expected, filtered)
				}
			default:
				t.Fatalf("Unsupported type %T", tt.input)
			}
		})
	}
}
