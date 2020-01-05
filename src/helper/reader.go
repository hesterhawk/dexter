package helper

import (
  "bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile(path string) {

  file, err := os.Open(path)
  CheckError(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	for _, eachline := range lines {
		fmt.Println("line: " + eachline)
	}
}

func CheckError(e error) {
  if e != nil {
    log.Fatalf("[!!] failed opening file: %s", e)
  }
}
