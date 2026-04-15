package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"mixed numbers", -2, 5, 3},
		{"zeros", 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name          string
		a, b          float64
		expected      float64
		expectedError bool
	}{
		{"normal division", 10.0, 2.0, 5.0, false},
		{"division with fraction", 5.0, 2.0, 2.5, false},
		{"division by zero", 10.0, 0.0, 0.0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)

			if (err != nil) != tt.expectedError {
				t.Errorf("Divide(%f, %f) error = %v, expectedError %v", tt.a, tt.b, err, tt.expectedError)
				return
			}

			if !tt.expectedError && result != tt.expected {
				t.Errorf("Divide(%f, %f) = %f; expected %f", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}
