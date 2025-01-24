package evaluator

import (
	"fmt"
	"math"
	"strings"

	"github.com/Knetic/govaluate"
)

// variables stores user-defined variables
var variables = make(map[string]interface{})

// functions maps custom functions to their implementations
var functions = map[string]govaluate.ExpressionFunction{
	"sqrt": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("sqrt expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("sqrt expects a numeric argument")
		}
		if val < 0 {
			return nil, fmt.Errorf("square root of negative number")
		}
		return math.Sqrt(val), nil
	},
	"sin": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("sin expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("sin expects a numeric argument")
		}
		return math.Sin(val), nil
	},
	"cos": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("cos expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("cos expects a numeric argument")
		}
		return math.Cos(val), nil
	},
	"tan": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("tan expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("tan expects a numeric argument")
		}
		return math.Tan(val), nil
	},
	"fact": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("fact expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("fact expects a numeric argument")
		}
		if val < 0 {
			return nil, fmt.Errorf("factorial is not defined for negative numbers")
		}
		result := 1.0
		for i := 1.0; i <= val; i++ {
			result *= i
		}
		return result, nil
	},
	"log": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("log expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("log expects a numeric argument")
		}
		return math.Log(val), nil
	},
	"log10": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("log10 expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("log10 expects a numeric argument")
		}
		return math.Log10(val), nil
	},
	"exp": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("exp expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("exp expects a numeric argument")
		}
		return math.Exp(val), nil
	},
	"pow": func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("pow expects exactly 2 arguments")
		}
		base, ok1 := args[0].(float64)
		exponent, ok2 := args[1].(float64)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("pow expects numeric arguments")
		}
		return math.Pow(base, exponent), nil
	},
	"abs": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("abs expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("abs expects a numeric argument")
		}
		return math.Abs(val), nil
	},
	"ceil": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("ceil expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("ceil expects a numeric argument")
		}
		return math.Ceil(val), nil
	},
	"floor": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("floor expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("floor expects a numeric argument")
		}
		return math.Floor(val), nil
	},
	"round": func(args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, fmt.Errorf("round expects exactly 1 argument")
		}
		val, ok := args[0].(float64)
		if !ok {
			return nil, fmt.Errorf("round expects a numeric argument")
		}
		return math.Round(val), nil
	},
	"min": func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("min expects exactly 2 arguments")
		}
		val1, ok1 := args[0].(float64)
		val2, ok2 := args[1].(float64)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("min expects numeric arguments")
		}
		return math.Min(val1, val2), nil
	},
	"max": func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, fmt.Errorf("max expects exactly 2 arguments")
		}
		val1, ok1 := args[0].(float64)
		val2, ok2 := args[1].(float64)
		if !ok1 || !ok2 {
			return nil, fmt.Errorf("max expects numeric arguments")
		}
		return math.Max(val1, val2), nil
	},
}

// Evaluate evaluates a mathematical expression or assigns a variable
func Evaluate(expression string) (interface{}, error) {
	// Check for empty expression
	if strings.TrimSpace(expression) == "" {
		return nil, fmt.Errorf("empty expression")
	}

	// Split the input into individual statements (e.g., "A = 5; A + 3")
	statements := strings.Split(expression, ";")
	var result interface{}

	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}

		// Check if the statement is a variable assignment (e.g., "A = 5")
		if strings.Contains(stmt, "=") {
			parts := strings.SplitN(stmt, "=", 2)
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid variable assignment: %s", stmt)
			}

			// Extract variable name and value
			varName := strings.TrimSpace(parts[0])
			varValue := strings.TrimSpace(parts[1])

			// Evaluate the value expression
			expr, err := govaluate.NewEvaluableExpressionWithFunctions(varValue, functions)
			if err != nil {
				return nil, fmt.Errorf("invalid value expression: %v", err)
			}

			val, err := expr.Evaluate(variables)
			if err != nil {
				return nil, fmt.Errorf("failed to evaluate value expression: %v", err)
			}

			// Store the variable in the map
			variables[varName] = val
			continue
		}

		// Replace the exponent operator (^) with ** (supported by govaluate)
		stmt = strings.ReplaceAll(stmt, "^", "**")

		// Evaluate the expression using the stored variables and custom functions
		expr, err := govaluate.NewEvaluableExpressionWithFunctions(stmt, functions)
		if err != nil {
			return nil, fmt.Errorf("invalid expression: %v", err)
		}

		result, err = expr.Evaluate(variables)
		if err != nil {
			return nil, fmt.Errorf("failed to evaluate expression: %v", err)
		}

		// Check for division by zero
		if strings.Contains(stmt, "/") {
			if val, ok := result.(float64); ok && math.IsInf(val, 0) {
				return nil, fmt.Errorf("division by zero")
			}
		}
	}

	return result, nil
}
