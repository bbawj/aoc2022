package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	priority := 0
	for scanner.Scan() {
		line := scanner.Text()
		n := len(line)
		first := line[:n/2]
		second := line[n/2:]
		seen := make(map[rune]bool)
		for _, c := range first {
			if _, found := seen[c]; !found {
				seen[c] = true
			}
		}
		for _, c := range second {
			if _, found := seen[c]; found {
				if int(c) < 'a' {
					priority += int(c) - 'A' + 27
				} else {
					priority += int(c) - 'a' + 1
				}
				break
			}
		}
	}

	fmt.Println(priority)
}
