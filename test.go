package main

import (
	//"fmt"
	"io/ioutil"
	"strings"
	"os"
)

func main() {
	var path = "/home/mati/go/dexter/bin"

	dir, _ := ioutil.ReadDir(path)

	for _, d := range dir {

		var file = strings.Split(d.Name(), ".")

		if file[1] == "html" {
			os.RemoveAll(path + "/" + d.Name())
		}
	}
}