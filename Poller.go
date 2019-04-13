package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

var configDir string

func init() {
	flag.StringVar(&configDir, "config", "/etc/SnmpPoller", "Config path")
	flag.Parse()
}

func main() {

	fmt.Println(configDir)
	ImportConfig(configDir)

	fmt.Println(viper.Get("Output.Number"))
}
