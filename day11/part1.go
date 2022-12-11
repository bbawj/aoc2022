package main

import (
	"bbawj/aoc2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Monkey struct {
	Items      []int64
	Operation  string
	Arg        int64
	Test       int64
	PassMonkey int64
	FailMonkey int64
	Count      int64
}

func main() {
	utils.GetInput(11)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// line := scanner.Text()
		// fmt.Println(line)
	}
	m0 := &Monkey{
		Items:      []int64{83, 97, 95, 67},
		Operation:  "*",
		Arg:        19,
		Test:       17,
		PassMonkey: 2,
		FailMonkey: 7,
	}
	m1 := &Monkey{
		Items:      []int64{71, 70, 79, 88, 56, 70},
		Operation:  "+",
		Arg:        2,
		Test:       19,
		PassMonkey: 7,
		FailMonkey: 0,
	}
	m2 := &Monkey{
		Items:      []int64{98, 51, 51, 63, 80, 85, 84, 95},
		Operation:  "+",
		Arg:        7,
		Test:       7,
		PassMonkey: 4,
		FailMonkey: 3,
	}
	m3 := &Monkey{
		Items:      []int64{77, 90, 82, 80, 79},
		Operation:  "+",
		Arg:        1,
		Test:       11,
		PassMonkey: 6,
		FailMonkey: 4,
	}
	m4 := &Monkey{
		Items:      []int64{68},
		Operation:  "*",
		Arg:        5,
		Test:       13,
		PassMonkey: 6,
		FailMonkey: 5,
	}
	m5 := &Monkey{
		Items:      []int64{60, 94},
		Operation:  "+",
		Arg:        5,
		Test:       3,
		PassMonkey: 1,
		FailMonkey: 0,
	}
	m6 := &Monkey{
		Items:      []int64{81, 51, 85},
		Operation:  "*",
		Arg:        0,
		Test:       5,
		PassMonkey: 5,
		FailMonkey: 1,
	}
	m7 := &Monkey{
		Items:      []int64{98, 81, 63, 65, 84, 71, 84},
		Operation:  "+",
		Arg:        3,
		Test:       2,
		PassMonkey: 2,
		FailMonkey: 3,
	}
	divisor := int64(1)
	m := []*Monkey{m0, m1, m2, m3, m4, m5, m6, m7}
	for _, monkey := range m {
		divisor *= monkey.Test
	}
	for i := 0; i < 10000; i++ {
		for _, monkey := range m {
			for _, item := range monkey.Items {
				i := &item
				monkey.Count++
				if monkey.Operation == "*" {
					if monkey.Arg == 0 {
						*i *= *i
					} else {
						*i *= monkey.Arg
					}
				} else {
					*i += monkey.Arg
				}

				*i %= divisor
				// *i = int64(math.Floor(float64(*i) / 3))

				if *i%monkey.Test == 0 {
					m[monkey.PassMonkey].Items = append(m[monkey.PassMonkey].Items, *i)
				} else {
					m[monkey.FailMonkey].Items = append(m[monkey.FailMonkey].Items, *i)
				}
			}
			monkey.Items = []int64{}
		}
	}
	sort.Slice(m, func(i, j int) bool {
		return m[i].Count > m[j].Count
	})

	fmt.Println(m[0].Count * m[1].Count)
}
