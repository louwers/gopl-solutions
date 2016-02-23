/*
This is my solution to exercise 1.4 of "The Go Programming Language".
It was my goal to only use concepts that have been explained in the book thus far.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
		// Printing
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	} else {
		counts := make(map[string]map[int]int)
		for fileIdx, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			fileCounts := make(map[string]int)
			countLines(f, fileCounts)
			for line, amount := range fileCounts {
				if counts[line] == nil {
					counts[line] = make(map[int]int)
				}
				counts[line][0] += amount // keeps track of the total
                :x
			}
			f.Close()
		}
		// Printing
		for line, arr := range counts {
			if arr[0] > 1 {
				fmt.Printf("%s\n", line)
			} else {
				continue
			}
			for fileNo, count := range arr {
				if fileNo == 0 { // skip the total count
					continue
				}
				if count > 0 {
					fmt.Printf("- %s\n", os.Args[fileNo])
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
