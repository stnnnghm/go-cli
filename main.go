package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Use Open() func from 'os' pkg,
	// which returns File Pointer to perform I/O
	// operations on File
	f, err := os.Open("sample.log")
	if err != nil {
		fmt.Println("error opening log file: %+v", err)
		return
	}

	// Ensure file is closed before exiting
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

		// Setup filtering log for log-level [ERROR]
		if strings.Contains(s, "[ERROR]") {
			fmt.Println(s)
		}
	}
}
