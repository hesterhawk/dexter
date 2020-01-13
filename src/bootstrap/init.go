package bootstrap

import (	
	"io/ioutil"
	"strings"
	"helper"
	"os"
)

var config helper.Config

func Init() {
	/*
		Load main configuration
	*/
	config = helper.LoadConfig()

	/*
		Clean Bin directory
	*/
	CleanDirFromHtml(config.AppPath + "/bin")
}

func Config() helper.Config {

	return config
}

func CleanDirFromHtml(path string) {

	dir, _ := ioutil.ReadDir(path)

	for _, d := range dir {

		var file = strings.Split(d.Name(), ".")

		if 2 == len(file) && file[1] == "html" {
			os.RemoveAll(path + "/" + d.Name())
		}
	}
}