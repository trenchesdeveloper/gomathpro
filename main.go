/*
Copyright © 2025 Opeyemi Samuel <opeyemisamuel222@gmail.com>
*/
package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/trenchesdeveloper/gomathpro/cmd"
)

var log = logrus.New()

func main() {
	// Initialize logging
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)

	// Execute the root command
	if err := cmd.Execute(); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Command execution failed")
		os.Exit(1)
	}
}