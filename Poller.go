package main

import (
	"flag"
	"github.com/sirupsen/logrus"
)

var configDir string
var logger *logrus.Logger

func init() {
	flag.StringVar(&configDir, "config", "/etc/SnmpPoller", "Config path")
	flag.Parse()

}

func main() {
	//
	// Import configuration
	//
	ImportConfig(configDir)

	logger = SetupLogger()

	logger.Info("Configuration loaded")
}
