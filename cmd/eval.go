package cmd

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/trenchesdeveloper/gomathpro/internal/evaluator"
)

var log = logrus.New()

// evalCmd represents the eval command
var evalCmd = &cobra.Command{
	Use:   "eval [expression]",
	Short: "Evaluate a mathematical expression",
	Long:  `Evaluate a mathematical expression with support for variables, exponents, factorials, and functions. Example: A = 5; B = 7; A + B`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		expression := args[0]

		// Split the input into individual statements (e.g., "A = 5; B = 7; A + B")
		statements := strings.Split(expression, ";")

		// Evaluate each statement
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			result, err := evaluator.Evaluate(stmt)
			if err != nil {
				log.WithFields(logrus.Fields{
					"error":      err,
					"expression": stmt,
				}).Error("Failed to evaluate expression")
				fmt.Printf("Error: %v\n", err)
				return
			}

			// If the result is not nil, print it (e.g., for expressions like "A + B")
			if result != nil {
				log.WithFields(logrus.Fields{
					"expression": stmt,
					"result":     result,
				}).Info("Expression evaluated successfully")
				fmt.Printf("Result: %v\n", result)
			}
		}
	},
}

func init() {
	// Add the eval command to the root command
	RootCmd.AddCommand(evalCmd)
}