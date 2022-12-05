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

type Stack []string

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Pop() string {
	if s.Len() == 0 {
		return ""
	} else {
		el := (*s)[0]
		*s = (*s)[1:]
		return el
	}
}

func (s *Stack) Push(val string) {
	*s = append([]string{val}, *s...)
}

func (s *Stack) Print() {
	var sb strings.Builder
	for _, v := range *s {
		sb.WriteString(v)
	}
	fmt.Println(sb.String())
}

func main() {
	utils.GetInput(5)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	crates := map[int64]*Stack{
		1: {"F", "L", "M", "W"},
		2: {"F", "M", "V", "Z", "B"},
		3: {"Q", "L", "S", "R", "V", "H"},
		4: {"J", "T", "M", "P", "Q", "V", "S", "F"},
		5: {"W", "S", "L"},
		6: {"W", "J", "R", "M", "P", "V", "F"},
		7: {"F", "R", "N", "P", "C", "Q", "J"},
		8: {"B", "R", "W", "Z", "S", "P", "H", "V"},
		9: {"W", "Z", "H", "G", "C", "J", "M", "B"},
	}
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, "move") {
			continue
		}

		val := strings.Split(line, " ")
		amt, _ := strconv.ParseInt(val[1], 10, 64)
		from, _ := strconv.ParseInt(val[3], 10, 64)
		to, _ := strconv.ParseInt(val[5], 10, 64)
		temp := &Stack{}
		for i := int64(0); i < amt; i++ {
			temp.Push(crates[from].Pop())
		}
		for i := int64(0); i < amt; i++ {
			crates[to].Push(temp.Pop())
		}
	}

	for k, v := range crates {
		fmt.Print(k, " ")
		v.Print()
	}
}
