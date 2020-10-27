package main

import (
	"github.com/netlify/gotrue/conf"
	"github.com/sirupsen/logrus"

	api "github.com/GitJournal/gotrue-go"
)

var configFile = ""

func execGoTrueWithConfig() {
	globalConfig, err := conf.LoadGlobal(configFile)
	if err != nil {
		logrus.Fatalf("Failed to load configuration: %+v", err)
	}
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		logrus.Fatalf("Failed to load configuration: %+v", err)
	}

	api := api.GoTrueAPI{}
	api.Serve(globalConfig, config)
}

func main() {
	execGoTrueWithConfig()
}
