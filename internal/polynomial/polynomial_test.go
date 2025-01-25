package polynomial_test

import (
	"math"
	"math/cmplx"
	"reflect"
	"testing"

	"github.com/trenchesdeveloper/gomathpro/internal/polynomial"
)

// A small epsilon for floating-point comparisons
const epsilon = 1e-9

// Helper: compare two slices of float64 with a tolerance.
func floatsAlmostEqual(a, b []float64, tol float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > tol {
			return false
		}
	}
	return true
}

// Helper: compare two slices of complex128 with a tolerance on real & imag parts.
func complexesAlmostEqual(a, b []complex128, tol float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if cmplx.Abs(a[i]-b[i]) > tol {
			return false
		}
	}
	return true
}

// ------------------------------------------------------------
// 1) TEST ParsePolynomial
// ------------------------------------------------------------
func TestParsePolynomial(t *testing.T) {
	tests := []struct {
		input   string
		want    []float64 // c0, c1, c2, ...
		wantErr bool
	}{
		// Constant only
		{"5", []float64{5}, false},
		{"-3", []float64{-3}, false},
		// Simple linear
		{"x", []float64{0, 1}, false},
		{"-x", []float64{0, -1}, false},
		{"2x", []float64{0, 2}, false},
		{"2x+1", []float64{1, 2}, false},
		// Simple quadratic
		{"x^2", []float64{0, 0, 1}, false},
		{"-3x^2", []float64{0, 0, -3}, false},
		{"x^2 - 5x + 6", []float64{6, -5, 1}, false},
		// Mixed degrees (cubic example)
		{"x^3 - 6x^2 + 11x - 6", []float64{-6, 11, -6, 1}, false},
		// Invalid
		{"", nil, true},
		{"abc", nil, true},
		{"x^", nil, true},
		{"^2", nil, true},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			got, err := polynomial.ParsePolynomial(tc.input)
			if (err != nil) != tc.wantErr {
				t.Fatalf("ParsePolynomial(%q) error = %v, wantErr = %v", tc.input, err, tc.wantErr)
			}
			if err == nil && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ParsePolynomial(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

// ------------------------------------------------------------
// 2) TEST FindRoots
// ------------------------------------------------------------
func TestFindRoots(t *testing.T) {
	tests := []struct {
		name      string
		coeffs    []float64 // c0, c1, c2, c3,...
		wantRoots []complex128
		wantErr   bool
	}{
		// Linear: 2x - 4 = 0 => x=2
		{
			name:      "Linear: 2x-4 => x=2",
			coeffs:    []float64{-4, 2},
			wantRoots: []complex128{complex(2, 0)},
			wantErr:   false,
		},
		// Linear: -3x + 6 => x=-2
		{
			name:      "Linear: -3x+6 => x=2",
			coeffs:    []float64{6, -3},
			wantRoots: []complex128{complex(2, 0)},
			wantErr:   false,
		},
		// Quadratic: x^2 - 5x + 6 => c0=6,c1=-5,c2=1 => roots:2,3
		{
			name:      "Quadratic: x^2-5x+6 => 2,3",
			coeffs:    []float64{6, -5, 1},
			wantRoots: []complex128{complex(2, 0), complex(3, 0)},
			wantErr:   false,
		},
		// Quadratic with complex roots: x^2+1=0 => c0=1, c1=0, c2=1 => roots: i, -i
		{
			name:      "Quadratic: x^2+1 => ±i",
			coeffs:    []float64{1, 0, 1},
			wantRoots: []complex128{complex(0, 1), complex(0, -1)},
			wantErr:   false,
		},
		// Higher-degree: x^3 - 6x^2 + 11x - 6 => 1,2,3
		{
			name:      "Cubic: x^3-6x^2+11x-6 => 1,2,3",
			coeffs:    []float64{-6, 11, -6, 1},
			wantRoots: []complex128{complex(1, 0), complex(2, 0), complex(3, 0)},
			wantErr:   false,
		},
		// Edge case: no coeffs
		{
			name:      "Empty coeffs => error",
			coeffs:    []float64{},
			wantRoots: nil,
			wantErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			roots, err := polynomial.FindRoots(tc.coeffs)
			if (err != nil) != tc.wantErr {
				t.Fatalf("FindRoots(%v) error = %v, wantErr %v", tc.coeffs, err, tc.wantErr)
			}
			if err == nil {
				// Because Durand-Kerner might list the same roots in a different order or with small numeric noise,
				// we do a "set" or "multiset" comparison with a tolerance
				if len(roots) != len(tc.wantRoots) {
					t.Errorf("FindRoots(%v) gave %v roots, want %v",
						tc.coeffs, len(roots), len(tc.wantRoots))
				}
				// We'll do a simpler approach: sort them by real part, then compare
				// (Also round imaginary parts for numeric stabilities if needed.)
				// For test brevity, let's do direct or a tolerance check:
				// In practice, you might want to do more robust matching (like a pairwise match).
				if !matchRootsWithTolerance(roots, tc.wantRoots, 1e-6) {
					t.Errorf("FindRoots(%v) = %v, want approx. %v", tc.coeffs, roots, tc.wantRoots)
				}
			}
		})
	}
}

// matchRootsWithTolerance tries to match each root in 'got' to 'want'
// allowing for small floating inaccuracies, ignoring permutation.
func matchRootsWithTolerance(got, want []complex128, tol float64) bool {
	if len(got) != len(want) {
		return false
	}
	used := make([]bool, len(want)) // track which want-root has been matched
	for _, g := range got {
		found := false
		for j, w := range want {
			if !used[j] && cmplx.Abs(g-w) < tol {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// ------------------------------------------------------------
// 3) TEST Factorize
// ------------------------------------------------------------
func TestFactorize(t *testing.T) {
	tests := []struct {
		name    string
		coeffs  []float64
		want    []string
		wantErr bool
	}{
		{
			name:   "Linear: 2x - 4 => factor (x - 2)",
			coeffs: []float64{-4, 2},
			want:   []string{"(x - 2.00)"},
		},
		{
			name:   "Linear: x + 1 => factor (x - -1.00)",
			coeffs: []float64{1, 1},
			want:   []string{"(x - -1.00)"},
		},
		{
			name:   "Quadratic: x^2 - 5x + 6 => (x - 2.00)(x - 3.00)",
			coeffs: []float64{6, -5, 1},
			want:   []string{"(x - 2.00)", "(x - 3.00)"},
		},
		{
			name:    "Quadratic complex => error",
			coeffs:  []float64{1, 0, 1}, // x^2+1=0 => i, -i => not factorable over reals
			wantErr: true,
		},
		{
			name:    "Cubic => error not supported",
			coeffs:  []float64{-6, 11, -6, 1}, // x^3 - 6x^2 + ...
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			factors, err := polynomial.Factorize(tc.coeffs)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Factorize(%v) error = %v, wantErr = %v", tc.coeffs, err, tc.wantErr)
			}
			if err == nil && !reflect.DeepEqual(factors, tc.want) {
				t.Errorf("Factorize(%v) = %v, want %v", tc.coeffs, factors, tc.want)
			}
		})
	}
}

// ------------------------------------------------------------
// 4) TEST Interpolate
// ------------------------------------------------------------
func TestInterpolate(t *testing.T) {
	t.Run("Simple Interpolation: points on a line y=2x+1", func(t *testing.T) {
		// y = 2x + 1 => c0=1, c1=2
		points := [][2]float64{
			{0, 1}, // y(0)=1
			{1, 3}, // y(1)=3
		}
		coeffs, err := polynomial.Interpolate(points)
		if err != nil {
			t.Fatalf("Interpolate error: %v", err)
		}
		// Expect 2 coefficients: c0=1, c1=2
		want := []float64{1, 2}
		if !floatsAlmostEqual(coeffs, want, epsilon) {
			t.Errorf("Interpolate(%v) = %v, want %v", points, coeffs, want)
		}
	})

	t.Run("Quadratic Interpolation: points on y=x^2", func(t *testing.T) {
		// If we pick x=0,1,2 => y= 0,1,4
		points := [][2]float64{
			{0, 0}, // y(0)=0
			{1, 1}, // y(1)=1
			{2, 4}, // y(2)=4
		}
		coeffs, err := polynomial.Interpolate(points)
		if err != nil {
			t.Fatalf("Interpolate error: %v", err)
		}
		// Expect c0=0, c1=0, c2=1, i.e. y= x^2
		want := []float64{0, 0, 1}
		if !floatsAlmostEqual(coeffs, want, epsilon) {
			t.Errorf("Interpolate(%v) = %v, want %v", points, coeffs, want)
		}
	})

	t.Run("No Points => error", func(t *testing.T) {
		points := [][2]float64{}
		_, err := polynomial.Interpolate(points)
		if err == nil {
			t.Errorf("Interpolate(%v) expected error, got nil", points)
		}
	})
}
