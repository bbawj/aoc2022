package main

import (
	"bbawj/aoc2022/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	utils.GetInput(13)
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	packets := make([]any, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if line == "" {
			continue
		} else {
			a, _ := parseExpression(line)
			packets = append(packets, a)
		}
	}
	div1, _ := parseExpression("[[2]]")
	div2, _ := parseExpression("[[6]]")
	packets = append(packets, div1)
	packets = append(packets, div2)
	sort.Slice(packets, func(i, j int) bool {
		return less(packets[i], packets[j]) > 0
	})
	var start, end int
	for i := 0; i < len(packets); i++ {
		str := fmt.Sprint(packets[i])
		if str == "[[[2]]]" {
			start = i + 1
		}
		if str == "[[[6]]]" {
			end = i + 1
		}
	}
	fmt.Println(start * end)
}

func less(a, b any) int {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt {
		if bInt < aInt {
			return -1
		} else if bInt > aInt {
			return 1
		}
		return 0
	}
	aArr, aIsArr := a.([]any)
	bArr, bIsArr := b.([]any)
	if !aIsArr {
		aArr = []any{aInt}
	}
	if !bIsArr {
		bArr = []any{bInt}
	}
	if len(aArr) == 0 && len(bArr) == 0 {
		return 0
	} else if len(aArr) == 0 && len(bArr) > 0 {
		return 1
	} else if len(aArr) > 0 && len(bArr) == 0 {
		return -1
	}
	for i := 0; i < len(aArr); i++ {
		if i >= len(bArr) {
			return -1
		}
		if sub := less(aArr[i], bArr[i]); sub != 0 {
			return sub
		}
	}

	if len(aArr) < len(bArr) {
		return 1
	} else if len(aArr) == len(bArr) {
		return less(aArr[len(aArr)-1], bArr[len(bArr)-1])
	}
	return 0
}

func parseExpression(expr string) (any, int) {
	chars := []rune(expr)
	idx := 0
	out := make([]any, 0)
	nChars := make([]rune, 0)
	for idx < len(chars) {
		char := chars[idx]
		switch char {
		case '[':
			x, i := parseExpression(string(chars[idx+1:]))
			out = append(out, x)
			idx += i + 1
		case ']':
			if len(nChars) > 0 {
				n, _ := strconv.Atoi(string(nChars))
				out = append(out, n)
				nChars = make([]rune, 0)
			}
			idx++
			return out, idx
		case ',':
			if len(nChars) > 0 {
				n, _ := strconv.Atoi(string(nChars))
				out = append(out, n)
				nChars = make([]rune, 0)
			}
			idx++
		default:
			nChars = append(nChars, char)
			idx++
		}
	}
	return out, idx
}
