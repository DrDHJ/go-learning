package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, data := range files {
			f, err := os.Open(data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "test1_4: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for line, m := range counts {
			sum := 0
			for _, n := range m {
				sum += n
			}
			if sum > 1 {
				for file, n := range m {
					fmt.Printf("%d\t%d\t%s\t%s\n", sum, n, file, line)
				}
			}

		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][f.Name()]++
	}

}
