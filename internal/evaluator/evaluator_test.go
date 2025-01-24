package evaluator

import (
	"testing"
)

// TestEvaluate tests the Evaluate function with various expressions
func TestEvaluate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
		hasError bool
	}{
		// Basic arithmetic
		{"Addition", "5 + 3", 8.0, false},
		{"Subtraction", "10 - 4", 6.0, false},
		{"Multiplication", "6 * 7", 42.0, false},
		{"Division", "20 / 4", 5.0, false},
		{"Division by zero", "10 / 0", nil, true},

		// Exponents
		{"Exponent", "2 ^ 3", 8.0, false},

		// Factorials
		{"Factorial", "fact(5)", 120.0, false},
		{"Factorial of negative", "fact(-1)", nil, true},

		// Functions
		{"Square root", "sqrt(16)", 4.0, false},
		{"Square root of negative", "sqrt(-1)", nil, true},
		{"Sine", "sin(0)", 0.0, false},
		{"Cosine", "cos(0)", 1.0, false},
		{"Tangent", "tan(0)", 0.0, false},
		{"Natural logarithm", "log(10)", 2.302585092994046, false},
		{"Base-10 logarithm", "log10(100)", 2.0, false},
		{"Exponential", "exp(2)", 7.38905609893065, false},
		{"Power", "pow(2, 3)", 8.0, false},

		// Absolute value, ceiling, floor, round
		{"Absolute value", "abs(-5)", 5.0, false},
		{"Ceiling", "ceil(3.2)", 4.0, false},
		{"Floor", "floor(3.8)", 3.0, false},
		{"Round", "round(3.5)", 4.0, false},

		// Min and max
		{"Minimum", "min(5, 10)", 5.0, false},
		{"Maximum", "max(5, 10)", 10.0, false},

		// Variables
		{"Variable assignment", "A = 5; A + 3", 8.0, false},
		{"Variable reuse", "B = 7; B * 2", 14.0, false},
		{"Undefined variable", "C + 5", nil, true},

		// Invalid expressions
		{"Invalid expression", "5 + * 3", nil, true},
		{"Empty expression", "", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Evaluate(tt.input)

			// Check for errors
			if tt.hasError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
				return
			}

			// Check for unexpected errors
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Compare results
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestFunctions tests individual functions in the evaluator
func TestFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function string
		args     []interface{}
		expected interface{}
		hasError bool
	}{
		{"Square root", "sqrt", []interface{}{16.0}, 4.0, false},
		{"Square root of negative", "sqrt", []interface{}{-1.0}, nil, true},
		{"Factorial", "fact", []interface{}{5.0}, 120.0, false},
		{"Factorial of negative", "fact", []interface{}{-1.0}, nil, true},
		{"Natural logarithm", "log", []interface{}{10.0}, 2.302585092994046, false},
		{"Base-10 logarithm", "log10", []interface{}{100.0}, 2.0, false},
		{"Exponential", "exp", []interface{}{2.0}, 7.38905609893065, false},
		{"Power", "pow", []interface{}{2.0, 3.0}, 8.0, false},
		{"Absolute value", "abs", []interface{}{-5.0}, 5.0, false},
		{"Ceiling", "ceil", []interface{}{3.2}, 4.0, false},
		{"Floor", "floor", []interface{}{3.8}, 3.0, false},
		{"Round", "round", []interface{}{3.5}, 4.0, false},
		{"Minimum", "min", []interface{}{5.0, 10.0}, 5.0, false},
		{"Maximum", "max", []interface{}{5.0, 10.0}, 10.0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, ok := functions[tt.function]
			if !ok {
				t.Errorf("Function %s not found", tt.function)
				return
			}

			result, err := fn(tt.args...)

			// Check for errors
			if tt.hasError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
				return
			}

			// Check for unexpected errors
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Compare results
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}