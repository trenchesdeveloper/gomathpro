/*
Copyright © 2025 Opeyemi Samuel <opeyemisamuel222@gmail.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command
var RootCmd = &cobra.Command{
	Use:   "gomathpro",
	Short: "A CLI tool for mathematical computations",
	Long:  `A CLI tool for performing mathematical computations like arithmetic, algebra, calculus, and more.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return RootCmd.Execute()
}