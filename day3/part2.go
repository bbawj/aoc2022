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
		freq := make(map[rune]int)
		for i := 1; i <= 3; i++ {
			seen := make(map[rune]bool)
			line := scanner.Text()
			fmt.Println(line)
			if i != 3 {
				scanner.Scan()
			}
			for _, c := range line {
				if _, found := freq[c]; !found {
					freq[c] = 1
					seen[c] = true
				} else if found && !seen[c] {
					freq[c] += 1
					seen[c] = true
					if freq[c] == 3 {
						// fmt.Println(string(c))
						if int(c) < 'a' {
							priority += int(c) - 'A' + 27
						} else {
							priority += int(c) - 'a' + 1
						}
						break
					}
				}
			}
		}
	}

	fmt.Println(priority)
}
