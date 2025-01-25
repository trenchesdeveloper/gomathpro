package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/trenchesdeveloper/gomathpro/internal/polynomial"
)

// polynomialCmd represents the polynomial command
var polynomialCmd = &cobra.Command{
	Use:   "polynomial",
	Short: "Perform polynomial operations",
	Long:  `Perform polynomial operations like finding roots, factorization, and interpolation.`,
}

// rootsCmd represents the roots command
var rootsCmd = &cobra.Command{
	Use:   "roots [coefficients]",
	Short: "Find the roots of a polynomial",
	Long:  `Find the roots of a polynomial given its coefficients. Example: gomathpro polynomial roots 1 -3 2`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		polyStr := strings.Join(args, "")
		coefficients, err := polynomial.ParsePolynomial(polyStr)
		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("Failed to parse polynomial")
			fmt.Printf("Error: %v\n", err)
			return
		}

		roots, err := polynomial.FindRoots(coefficients)
		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("Failed to find roots")
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Roots:")
		for _, root := range roots {
			fmt.Printf("- %v\n", root)
		}
	},
}

// factorizeCmd represents the factorize command
var factorizeCmd = &cobra.Command{
	Use:   "factorize [coefficients]",
	Short: "Factorize a polynomial",
	Long:  `Factorize a polynomial given its coefficients. Example: gomathpro polynomial factorize 1 -3 2`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		polyStr := strings.Join(args, "")
		coefficients, err := polynomial.ParsePolynomial(polyStr)
		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("Failed to parse polynomial")
			fmt.Printf("Error: %v\n", err)
			return
		}

		factors, err := polynomial.Factorize(coefficients)
		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("Failed to factorize polynomial")
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Factors:")
		for _, factor := range factors {
			fmt.Printf("- %s\n", factor)
		}
	},
}

// interpolateCmd represents the interpolate command
var interpolateCmd = &cobra.Command{
	Use:   "interpolate [x1 y1 x2 y2 ...]",
	Short: "Interpolate a polynomial",
	Long:  `Interpolate a polynomial given a set of points. Example: gomathpro polynomial interpolate 1 2 3 4`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)%2 != 0 {
			log.Error("Invalid number of arguments. Expected pairs of x and y values.")
			fmt.Println("Error: Expected pairs of x and y values.")
			return
		}

		points := make([][2]float64, len(args)/2)
		for i := 0; i < len(args); i += 2 {
			x, err1 := strconv.ParseFloat(args[i], 64)
			y, err2 := strconv.ParseFloat(args[i+1], 64)
			if err1 != nil || err2 != nil {
				log.WithFields(logrus.Fields{
					"error": fmt.Errorf("invalid point: %s, %s", args[i], args[i+1]),
				}).Error("Invalid point")
				fmt.Printf("Error: Invalid point: %s, %s\n", args[i], args[i+1])
				return
			}
			points[i/2] = [2]float64{x, y}
		}

		coefficients, err := polynomial.Interpolate(points)
		if err != nil {
			log.WithFields(logrus.Fields{
				"error": err,
			}).Error("Failed to interpolate polynomial")
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Println("Interpolated Polynomial Coefficients:")
		for i, coeff := range coefficients {
			fmt.Printf("x^%d: %.2f\n", i, coeff)
		}
	},
}

func init() {
	// Add the polynomial command to the root command
	RootCmd.AddCommand(polynomialCmd)

	// Add child commands to the polynomial command
	polynomialCmd.AddCommand(rootsCmd)
	polynomialCmd.AddCommand(factorizeCmd)
	polynomialCmd.AddCommand(interpolateCmd)
}
