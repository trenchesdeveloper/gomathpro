package polynomial

import (
	"fmt"
	"math"
	"math/cmplx"
	"regexp"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

// ParsePolynomial parses a polynomial string into coefficients.
func ParsePolynomial(polyStr string) ([]float64, error) {
	// Remove all whitespace
	polyStr = strings.ReplaceAll(polyStr, " ", "")

	// Regex to match terms like "x^2", "-3x", "5", etc.
	termRegex := regexp.MustCompile(`([+-]?\d*\.?\d*x\^?\d*|[-+]?\d*\.?\d+)`)
	terms := termRegex.FindAllString(polyStr, -1)

	if len(terms) == 0 {
		return nil, fmt.Errorf("invalid polynomial format")
	}

	parsed := strings.Join(terms, "")
    if parsed != polyStr {
        return nil, fmt.Errorf("invalid polynomial format: extra/unmatched text in %q", polyStr)
    }


	// Find the highest degree in the polynomial
	maxDegree := 0
	for _, term := range terms {
		if strings.Contains(term, "x^") {
			parts := strings.Split(term, "^")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid term: %s", term)
			}
			degree, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid degree in term: %s", term)
			}
			if degree > maxDegree {
				maxDegree = degree
			}
		} else if strings.Contains(term, "x") {
			if maxDegree < 1 {
				maxDegree = 1
			}
		}
	}

	// Initialize coefficients slice
	coefficients := make([]float64, maxDegree+1)

	// Parse each term
	for _, term := range terms {
		if term == "" {
			continue
		}

		if strings.Contains(term, "x^") {
			// Handle terms like "x^2", "-3x^2", etc.
			parts := strings.Split(term, "x^")
			coeffStr := parts[0]
			if coeffStr == "" || coeffStr == "+" {
				coeffStr = "1"
			} else if coeffStr == "-" {
				coeffStr = "-1"
			}
			coeff, err := strconv.ParseFloat(coeffStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid coefficient in term: %s", term)
			}
			degree, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid degree in term: %s", term)
			}
			coefficients[degree] += coeff
		} else if strings.Contains(term, "x") {
			// Handle terms like "x", "-3x", etc.
			coeffStr := strings.TrimSuffix(term, "x")
			if coeffStr == "" || coeffStr == "+" {
				coeffStr = "1"
			} else if coeffStr == "-" {
				coeffStr = "-1"
			}
			coeff, err := strconv.ParseFloat(coeffStr, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid coefficient in term: %s", term)
			}
			coefficients[1] += coeff
		} else {
			// Handle constant terms like "5", "-3", etc.
			coeff, err := strconv.ParseFloat(term, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid constant term: %s", term)
			}
			coefficients[0] += coeff
		}
	}

	return coefficients, nil
}

// FindRoots finds the roots of a polynomial given its coefficients.
func FindRoots(coefficients []float64) ([]complex128, error) {
    if len(coefficients) == 0 {
        return nil, fmt.Errorf("no coefficients provided")
    }

    switch len(coefficients) {
    case 2:
        // Linear polynomial: c0 + c1*x = 0
        // => c1*x = -c0
        // => x = -c0/c1
        c0 := coefficients[0] // constant term
        c1 := coefficients[1] // x term
        if c1 == 0 {
            return nil, fmt.Errorf("invalid linear polynomial (coefficient of x cannot be zero)")
        }
        root := complex(-c0/c1, 0)
        return []complex128{root}, nil

    case 3:
        // Quadratic polynomial: c0 + c1*x + c2*x^2 = 0
        // => a = c2, b = c1, c = c0
        c0 := coefficients[0]
        c1 := coefficients[1]
        c2 := coefficients[2]
        if c2 == 0 {
            return nil, fmt.Errorf("invalid quadratic polynomial (coefficient of x^2 cannot be zero)")
        }
        discriminant := c1*c1 - 4*c2*c0

        if discriminant < 0 {
            realPart := -c1 / (2 * c2)
            imagPart := math.Sqrt(-discriminant) / (2 * c2)
            return []complex128{
                complex(realPart, imagPart),
                complex(realPart, -imagPart),
            }, nil
        } else {
            sqrtDisc := math.Sqrt(discriminant)
            root1 := (-c1 + sqrtDisc) / (2 * c2)
            root2 := (-c1 - sqrtDisc) / (2 * c2)
            return []complex128{
                complex(root1, 0),
                complex(root2, 0),
            }, nil
        }

    default:
        // For degree >= 3 polynomials (cubic, quartic, etc.), use Durand-Kerner
        return findRootsDurandKerner(coefficients), nil
    }
}

// findRootsDurandKerner finds the roots of a polynomial using the Durand-Kerner method.
func findRootsDurandKerner(coefficients []float64) []complex128 {
	n := len(coefficients) - 1 // Degree of the polynomial
	if n < 1 {
		return nil
	}

	// Initial guesses for roots
	roots := make([]complex128, n)
	for i := range roots {
		roots[i] = cmplx.Rect(1, 2*math.Pi*float64(i)/float64(n)) // Equally spaced on unit circle
	}

	// Durand-Kerner iteration
	for iter := 0; iter < 1000; iter++ {
		updated := make([]complex128, n)
		for i := range roots {
			numerator := evaluatePolynomial(coefficients, roots[i])
			denominator := complex(1, 0)
			for j := range roots {
				if i != j {
					denominator *= (roots[i] - roots[j])
				}
			}
			updated[i] = roots[i] - numerator/denominator
		}

		// Check for convergence
		converged := true
		for i := range roots {
			if cmplx.Abs(updated[i]-roots[i]) > 1e-10 {
				converged = false
				break
			}
		}
		if converged {
			break
		}

		roots = updated
	}

	return roots
}

// evaluatePolynomial evaluates a polynomial at a given complex point.
func evaluatePolynomial(coefficients []float64, x complex128) complex128 {
	result := complex(0, 0)
	for i, coeff := range coefficients {
		result += complex(coeff, 0) * cmplx.Pow(x, complex(float64(i), 0))
	}
	return result
}

// Factorize factorizes a polynomial into its irreducible factors.
func Factorize(coefficients []float64) ([]string, error) {
    if len(coefficients) == 0 {
        return nil, fmt.Errorf("no coefficients provided")
    }

    if len(coefficients) > 3 {
        return nil, fmt.Errorf("factorization is only supported for linear and quadratic polynomials")
    }

    switch len(coefficients) {
    case 2:
        // Linear: c0 + c1*x
        c0 := coefficients[0]
        c1 := coefficients[1]
        if c1 == 0 {
            return nil, fmt.Errorf("invalid linear polynomial (coefficient of x cannot be zero)")
        }
        // Root is -c0/c1, so factor is (x - root)
        root := -c0 / c1
        return []string{fmt.Sprintf("(x - %.2f)", root)}, nil

    case 3:
        // Quadratic: c0 + c1*x + c2*x^2
        c0 := coefficients[0]
        c1 := coefficients[1]
        c2 := coefficients[2]
        if c2 == 0 {
            return nil, fmt.Errorf("invalid quadratic polynomial (coefficient of x^2 cannot be zero)")
        }
        discriminant := c1*c1 - 4*c2*c0
        if discriminant < 0 {
            return nil, fmt.Errorf("cannot factorize polynomial with complex roots")
        }
        sqrtDisc := math.Sqrt(discriminant)
        r1 := (-c1 + sqrtDisc) / (2 * c2)
        r2 := (-c1 - sqrtDisc) / (2 * c2)

        // Ensure we return smaller root first
        if r1 > r2 {
            r1, r2 = r2, r1
        }

        return []string{
            fmt.Sprintf("(x - %.2f)", r1),
            fmt.Sprintf("(x - %.2f)", r2),
        }, nil

    default:
        return nil, fmt.Errorf("unsupported polynomial degree")
    }
}

// Interpolate interpolates a polynomial given a set of points.
func Interpolate(points [][2]float64) ([]float64, error) {
	if len(points) == 0 {
		return nil, fmt.Errorf("no points provided")
	}

	// Create a Vandermonde matrix and solve for coefficients
	n := len(points)
	v := mat.NewDense(n, n, nil)
	b := mat.NewVecDense(n, nil)

	for i := 0; i < n; i++ {
		x := points[i][0]
		y := points[i][1]
		for j := 0; j < n; j++ {
			v.Set(i, j, math.Pow(x, float64(j)))
		}
		b.SetVec(i, y)
	}

	var coeffs mat.VecDense
	err := coeffs.SolveVec(v, b)
	if err != nil {
		return nil, fmt.Errorf("failed to solve interpolation: %v", err)
	}

	coefficients := make([]float64, n)
	for i := 0; i < n; i++ {
		coefficients[i] = coeffs.AtVec(i)
	}

	return coefficients, nil
}
