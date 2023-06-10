package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <filename> <search_string>")
		os.Exit(1)
	}

	filename := os.Args[1]
	searchString := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchString) {
			fmt.Printf("Line %d: %s\n", lineNumber, line)
			found = true
		}
		lineNumber++
	}

	if !found {
		fmt.Printf("The search string '%s' was not found in the file '%s'\n", searchString, filename)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

