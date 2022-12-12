package main

import (
	"bbawj/aoc2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Queue []Pos

func (q *Queue) Push(val Pos) {
	*q = append(*q, val)
}

func (q *Queue) Pop() Pos {
	el := (*q)[0]
	*q = (*q)[1:]
	return el
}

func (q *Queue) Len() int {
	return len(*q)
}

func main() {
	utils.GetInput(12)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	grid := make([][]rune, 0)
	scanner := bufio.NewScanner(f)
	var end Pos
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		if strings.ContainsRune(line, 'S') {
			sIndex := strings.Index(line, "S")
			row[sIndex] = 'a'
		}
		if strings.ContainsRune(line, 'E') {
			sIndex := strings.Index(line, "E")
			end = Pos{sIndex, len(grid)}
			row[sIndex] = 'z'
		}
		// fmt.Println(line)
		grid = append(grid, row)
	}
	min := len(grid) * len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'a' {
				steps := bfs(grid, Pos{j, i}, end)
				if steps < min {
					min = steps
				}
			}
		}
	}
	fmt.Println(min)
}

func bfs(grid [][]rune, start, end Pos) int {
	moves := []Pos{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	visited := make(map[Pos]bool)
	visited[start] = true
	steps := -1
	q := make([]Queue, 0)
	temp := []Pos{start}
	q = append(q, temp)
	found := false
	for len(q) > 0 && !found {
		steps++
		level := q[0]
		q = q[1:]
		nextLevel := make(Queue, 0)
		for _, pos := range level {
			letter := grid[pos.y][pos.x]
			if pos.y == end.y && pos.x == end.x {
				found = true
				break
			}
			for _, move := range moves {
				nextPos := Pos{pos.x + move.x, pos.y + move.y}
				if nextPos.x < 0 || nextPos.x >= len(grid[0]) || nextPos.y < 0 || nextPos.y >= len(grid) {
					continue
				}
				if visited[nextPos] || int(grid[nextPos.y][nextPos.x])-int(letter) > 1 {
					continue
				}
				nextLevel.Push(nextPos)
				visited[nextPos] = true
			}
		}
		if nextLevel.Len() > 0 {
			q = append(q, nextLevel)
		}
	}
	if !found {
		return len(grid) * len(grid[0])
	}
	return steps
}
