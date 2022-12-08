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
	utils.GetInput(8)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		row := make([]int64, len(line))
		for i := range line {
			row[i], _ = strconv.ParseInt(string(line[i]), 10, 64)
		}
		grid = append(grid, row)
	}
  m := len(grid)
  n := len(grid[0])
  left := make([][]int64,0)
  right := make([][]int64,0)
  top := make([][]int64,0)
  bottom := make([][]int64,0)

  for i := 0; i < m; i++ {
    max := int64(-1)
    temp := make([]int64, n))
    for j := 0; j < n; j++ {
      temp[j] = max
      if grid[i][j] > max {
        max = grid[i][j]
      }
    }
    left = append(left, temp)
  }
}
