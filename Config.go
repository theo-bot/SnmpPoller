package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func ImportConfig(configPath string) {
	viper.SetConfigName("SnmpPoller")
	viper.SetConfigType("json")
	viper.AddConfigPath("/etc/SnmpPoller")
	viper.AddConfigPath("$HOME/.SnmpPoller")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}
}
