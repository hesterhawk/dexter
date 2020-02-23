package main

import (
	"os"
	"fmt"
	"time"
	"helper"
)

// Just run it..

func main() {

	argsWithProg := os.Args

	if len(argsWithProg) < 2 {
		fmt.Println("usage: go run new_draft.go <lang pl:en>")
		return
	}

	if argsWithProg[1] != "pl" && argsWithProg[1] != "en" {
		fmt.Println("usage: go run new_draft.go <lang pl:en>")
		return
	}

	var currentTime = time.Now()

	var config = helper.LoadConfig()

	var newDraft = config.AppPath + "/drafts/" + argsWithProg[1] + "/" + currentTime.Format("2006-01-02") + "-new.html"

	file, err := os.Create(newDraft)
	
	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("<!--\nTitle\nDescription\nCreated Date\nAuthor\nImage\nVisible:True\nUpdate:True\n!-->")
	
	fmt.Println("[+] New draft " + newDraft + " created!")
}