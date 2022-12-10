package main

import (
	"bbawj/aoc2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	utils.GetInput(10)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	cycleCount := 0
	register := int64(1)
	signal := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		val := strings.Split(line, " ")
		var arg int64
		op := val[0]
		if op == "addx" {
			arg, _ = strconv.ParseInt(val[1], 10, 64)
			for i := 0; i < 2; i++ {
				cycleCount++
				if (cycleCount-20)%40 == 0 {
					signal += cycleCount * int(register)
				}
			}
			register += arg
		} else {
			cycleCount++
			if (cycleCount-20)%40 == 0 {
				signal += cycleCount * int(register)
			}
		}
	}
	fmt.Println(signal)
}
