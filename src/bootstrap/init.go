package bootstrap

import "helper"

var config helper.Config

func Init() {
	config = helper.LoadConfig()
}

func Config() helper.Config {

	return config
}