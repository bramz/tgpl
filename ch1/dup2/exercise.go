package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filenames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, filenames)
			f.Close()
		}
	}
	for line, l := range filenames {
		if len(l) >= 2 {
			fmt.Printf("%v\t%s\n", l, line)
		}
	}
}

func checker(file string, data []string) bool {
	for _, d := range data {
		if file == d {
			return true
		}
	}
	return false
}

func countLines(f *os.File, filenames map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if !checker(f.Name(), filenames[input.Text()]) {
			filenames[input.Text()] = append(filenames[input.Text()], f.Name())
		}
	}
}
