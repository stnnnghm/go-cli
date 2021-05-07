package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	// Accept a value from the command line
	// We will accept a string

	// Arguments: 1. var name 2. default val 3. help desc
	path := flag.String("path", "sample.log", "Path to the log file")

	level := flag.String("level", "ERROR", "Log-level to filter. Options: [ ERROR, INFO, FATAL, DEBUG, WARN, TRACE ]")

	// Fetch values from command line using flag.Parse() and populate variable values declared above
	flag.Parse()

	// Pass the path's pointer to Open() since flag.String() returns a pointer
	f, err := os.Open(*path)
	if err != nil {
		fmt.Println("error opening log file: %+v", err)
		return
	}
	defer f.Close()


	// Use bufio pkg to read the contents of file
	// Reader is designed to read an input stream
	// In this case it is a file but readers can be anything
	// ie. stream from network conn, string, etc.
	r := bufio.NewReader(f)

	// Infinite for loop will read every line in a file
	for {
		// create reader that allows reading file contents
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}

		// Setup filtering log for log-level fetched from command line 
		// and stored in level var
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
